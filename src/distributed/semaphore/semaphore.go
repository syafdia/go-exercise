package semaphore

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	luaScriptReleaseSemaphore = `
		local keyResourceWithTimeout = KEYS[1]
		local keyResources = KEYS[2]
		local token = ARGV[1]

		local score = redis.call('ZSCORE', keyResourceWithTimeout, token)
		if not score or tonumber(score) == 0 then
			return 0
		end

		redis.call('LPUSH', keyResources, token)
		redis.call('ZREM', keyResourceWithTimeout, token)

		return 1
	`

	luaScriptCleanupDanglingResources = `
		local keyResourceWithTimeout = KEYS[1]
		local keyResources = KEYS[2]
		local minStr = ARGV[1]
		local maxStr = ARGV[2]

		local danglingResources = redis.call('ZRANGEBYSCORE', keyResourceWithTimeout, minStr, maxStr)
		if #danglingResources == 0 then
			return 1
		end

		redis.call('ZREM', keyResourceWithTimeout, unpack(danglingResources))
		redis.call('LPUSH', keyResources, unpack(danglingResources))

		return 1
	`
)

type Semaphore struct {
	maxConcurrency int
	name           string
	redisClient    *redis.Client
	acquireTimeout time.Duration
}

func New(
	maxConcurrency int,
	name string,
	redisClient *redis.Client,
	acquireTimeout time.Duration,
) (*Semaphore, error) {
	s := &Semaphore{
		maxConcurrency: maxConcurrency,
		name:           name,
		redisClient:    redisClient,
		acquireTimeout: acquireTimeout,
	}

	err := s.initResources()
	if err != nil {
		return nil, fmt.Errorf("semaphore: failed on init resources, %w", err)
	}

	return s, nil
}

// Acquire will borrow resource from the pool.
func (s *Semaphore) Acquire(ctx context.Context) (string, error) {
	err := s.cleanupDanglingResources(ctx)
	if err != nil {
		return "", fmt.Errorf("semaphore: failed when cleanup dangling resources, %w", err)
	}

	vals, err := s.redisClient.BRPop(ctx, s.acquireTimeout, s.keyResources()).Result()
	if err != nil {
		return "", fmt.Errorf("semaphore: failed when acquire, %w", err)
	}

	if len(vals) != 2 {
		return "", fmt.Errorf("semaphore: no resource available")
	}

	token := vals[1]

	member := &redis.Z{Score: s.timeNowUnix(), Member: token}
	_, err = s.redisClient.ZAdd(ctx, s.keyResourceWithTimeout(), member).Result()
	if err != nil {
		return "", fmt.Errorf("semaphore: failed when set timeout, %w", err)
	}

	return token, nil
}

// Release will return back borrowed resources to pool.
func (s *Semaphore) Release(ctx context.Context, token string) error {
	_, err := s.redisClient.
		Eval(ctx, luaScriptReleaseSemaphore, []string{s.keyResourceWithTimeout(), s.keyResources()}, token).
		Result()

	if err != nil {
		return fmt.Errorf("semaphore: failed when release, %w", err)
	}

	return nil
}

func (s *Semaphore) keyResources() string {
	return fmt.Sprintf("semaphore:resources:%s", s.name)
}

func (s *Semaphore) keyResourceWithTimeout() string {
	return fmt.Sprintf("semaphore:resource_timeout:%s", s.name)
}

func (s *Semaphore) timeNowUnix() float64 {
	return float64(time.Now().Unix())
}

func (s *Semaphore) initResources() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelFunc()

	tokens := s.generateTokens()
	_, err := s.redisClient.Pipelined(ctx, func(p redis.Pipeliner) error {
		_, err := p.Del(ctx, s.keyResources(), s.keyResourceWithTimeout()).Result()
		if err != nil {
			return err
		}

		_, err = p.LPush(ctx, s.keyResources(), tokens).Result()
		return err
	})

	return err
}

func (s *Semaphore) generateTokens() []string {
	tokens := []string{}
	for i := 0; i < s.maxConcurrency; i++ {
		tokens = append(tokens, uuid.New().String())
	}

	return tokens
}

func (s *Semaphore) cleanupDanglingResources(ctx context.Context) error {
	minStr := "-inf"
	maxStr := fmt.Sprintf("%d", int64(s.timeNowUnix()-s.acquireTimeout.Seconds()))

	_, err := s.redisClient.
		Eval(
			ctx,
			luaScriptCleanupDanglingResources,
			[]string{s.keyResourceWithTimeout(), s.keyResources()},
			minStr,
			maxStr,
		).
		Result()

	return err
}
