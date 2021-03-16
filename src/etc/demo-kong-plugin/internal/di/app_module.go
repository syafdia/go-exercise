package di

import (
	"sync"
)

type AppModule struct {
}

func NewAppModule() *AppModule {
	return &AppModule{}
}

var onceAppModule sync.Once
var appModule *AppModule

func GetAppModule() *AppModule {
	onceAppModule.Do(func() {
		appModule = NewAppModule()
	})

	return appModule
}
