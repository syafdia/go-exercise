package data

import (
	"context"
	"fmt"

	"github.com/syafdia/demo-kong-plugin/internal/domain/repo"
)

type tokenRepository struct {
}

func NewTokenRepository() repo.TokenRepository {
	return &tokenRepository{}
}

func (t *tokenRepository) ExchangeToken(ctx context.Context, key string) (string, error) {
	return fmt.Sprintf("%s-thisisasecrettoken", key), nil
}

func (t *tokenRepository) Invalidate(ctx context.Context, token string) error {
	return nil
}
