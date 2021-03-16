package delivery

import (
	"context"

	"github.com/Kong/go-pdk"
	"github.com/syafdia/demo-kong-plugin/internal/di"
	"github.com/syafdia/demo-kong-plugin/internal/domain/entity/keychecker"
)

// KeyCheckerConfig will hold all Kong configuration data
// and injected use case module.
type KeyCheckerConfig struct {
	APIKey        string `json:"api_key"`
	useCaseModule *di.UseCaseModule
}

// NewKeyChecker will create an instance of KeyCheckerConfig.
func NewKeyChecker() interface{} {
	return &KeyCheckerConfig{useCaseModule: di.GetUseCaseModule()}
}

func (conf *KeyCheckerConfig) Access(kong *pdk.PDK) {
	key, err := kong.Request.GetQueryArg("key")
	if err != nil {
		kong.Log.Err(err.Error())
	}

	ctx := context.Background()
	respHeaders := make(map[string][]string)
	respHeaders["Content-Type"] = append(respHeaders["Content-Type"], "application/json")

	err = conf.useCaseModule.ValidateKeyUseCase.Execute(ctx, keychecker.ValidateKeyInput{
		GivenKey: key,
		ValidKey: conf.APIKey,
	})
	if err != nil {
		switch err {
		case keychecker.ErrKeyEmpty, keychecker.ErrKeyNotValid:
			kong.Response.Exit(401, err.Error(), respHeaders)

		default:
			kong.Response.Exit(500, err.Error(), respHeaders)
		}

		return
	}

	token, err := conf.useCaseModule.GetTokenUseCase.Execute(ctx, key)
	if err != nil {
		kong.Response.Exit(500, err.Error(), respHeaders)
		return
	}

	kong.Response.SetHeader("X-Exchange-Token", token)
}
