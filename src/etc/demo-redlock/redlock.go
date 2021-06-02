package demoredlock

import (
	"context"
	"errors"
	"math/rand"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var (
	ErrAcquireResource = errors.New("locker: failed acquiring resource")
	ErrReleaseResource = errors.New("locker: failed releasing resource")
)

var scriptLock = `
if redis.call("EXISTS", KEYS[1]) == 1 then
	return 0
end

return redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
`

var scriptUnlock = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
else
	return 0
end
`

type Locker interface {
	Lock(ctx context.Context) error
	Unlock(ctx context.Context) error
}

type DLM struct {
	redisClients []*redis.Client
	quorum       int
	expiration   time.Duration
	drift        time.Duration
}

func NewDLM(redisClients []*redis.Client, expiration time.Duration, drift time.Duration) *DLM {
	return &DLM{
		redisClients: redisClients,
		expiration:   expiration,
		drift:        drift,
		quorum:       len(redisClients)/2 + 1,
	}
}

func (dlm *DLM) NewLocker(name string) Locker {
	return newLocker(dlm, name)
}

type locker struct {
	redisClients []*redis.Client
	expiration   time.Duration
	drift        time.Duration
	quorum       int
	name         string
	value        string
}

func newLocker(dlm *DLM, name string) Locker {
	return &locker{
		redisClients: dlm.redisClients,
		quorum:       dlm.quorum,
		name:         name,
		value:        generateRandomString(),
		expiration:   dlm.expiration,
		drift:        dlm.drift,
	}
}

func (l *locker) Lock(ctx context.Context) error {
	totalSuccess := 0

	for _, rc := range l.redisClients {
		start := time.Now()

		status, err := rc.Eval(ctx, scriptLock, []string{l.name}, l.value, l.expiration.Milliseconds()).Result()
		if err != nil {
			return err
		}

		ok := status == "OK"

		now := time.Now()
		isTimeValid := (l.expiration - (now.Sub(start) - l.drift)) > 0

		if ok && isTimeValid {
			totalSuccess++
		}
	}

	if totalSuccess < l.quorum {
		l.Unlock(ctx)
		return ErrAcquireResource
	}

	return nil
}

func (l *locker) Unlock(ctx context.Context) error {
	totalSuccess := 0

	for _, rc := range l.redisClients {

		status, err := rc.Eval(ctx, scriptUnlock, []string{l.name}, l.value).Result()
		if err != nil {
			return err
		}

		if status != int64(0) {
			totalSuccess++
		}
	}

	if totalSuccess < l.quorum {
		return ErrReleaseResource
	}

	return nil
}

func generateRandomString() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune,
		time.Now().Unix()%64)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
