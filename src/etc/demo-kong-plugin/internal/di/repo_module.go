package di

import (
	"sync"

	"github.com/syafdia/demo-kong-plugin/internal/data"
	"github.com/syafdia/demo-kong-plugin/internal/domain/repo"
)

type RepoModule struct {
	TokenRepository repo.TokenRepository
}

func NewRepoModule(appModule *AppModule) *RepoModule {
	return &RepoModule{
		TokenRepository: data.NewTokenRepository(),
	}
}

var onceRepoModule sync.Once
var repoModule *RepoModule

func GetRepoModule() *RepoModule {
	onceRepoModule.Do(func() {
		repoModule = NewRepoModule(GetAppModule())
	})

	return repoModule
}
