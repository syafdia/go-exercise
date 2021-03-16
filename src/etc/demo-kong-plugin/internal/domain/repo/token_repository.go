package repo

import "context"

type TokenRepository interface {
	ExchangeToken(ctx context.Context, key string) (string, error)
	Invalidate(ctx context.Context, token string) error
}
