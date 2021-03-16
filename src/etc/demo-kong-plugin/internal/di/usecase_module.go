package di

import (
	"sync"

	keycheckerusecase "github.com/syafdia/demo-kong-plugin/internal/domain/usecase/keychecker"
)

type UseCaseModule struct {
	ValidateKeyUseCase keycheckerusecase.ValidateKeyUseCase
	GetTokenUseCase    keycheckerusecase.GetTokenUseCase
}

func NewUseCaseModule(repoModule *RepoModule) *UseCaseModule {
	return &UseCaseModule{
		ValidateKeyUseCase: keycheckerusecase.NewValidateKeyUseCase(),
		GetTokenUseCase:    keycheckerusecase.NewGetTokenUseCase(repoModule.TokenRepository),
	}
}

var onceUseCaseModule sync.Once
var useCaseModule *UseCaseModule

func GetUseCaseModule() *UseCaseModule {
	onceUseCaseModule.Do(func() {
		useCaseModule = NewUseCaseModule(GetRepoModule())
	})

	return useCaseModule
}
