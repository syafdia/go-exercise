package keychecker

import (
	"context"
	"log"

	"github.com/syafdia/demo-kong-plugin/internal/domain/repo"
)

type GetTokenUseCase interface {
	Execute(ctx context.Context, key string) (string, error)
}

type getTokenUseCase struct {
	tokenRepo repo.TokenRepository
}

func NewGetTokenUseCase(tokenRepo repo.TokenRepository) GetTokenUseCase {
	return &getTokenUseCase{tokenRepo}
}

func (v *getTokenUseCase) Execute(ctx context.Context, key string) (string, error) {
	log.Println("[GetTokenUseCase] Start get token use case")
	return v.tokenRepo.ExchangeToken(ctx, key)
}
