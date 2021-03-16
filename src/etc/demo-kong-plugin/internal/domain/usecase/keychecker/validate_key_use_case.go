package keychecker

import (
	"context"
	"log"

	keycheckerentity "github.com/syafdia/demo-kong-plugin/internal/domain/entity/keychecker"
)

type ValidateKeyUseCase interface {
	Execute(ctx context.Context, forwardInput keycheckerentity.ValidateKeyInput) error
}

type validateKeyUseCase struct{}

func NewValidateKeyUseCase() ValidateKeyUseCase {
	return &validateKeyUseCase{}
}

func (v *validateKeyUseCase) Execute(ctx context.Context, forwardInput keycheckerentity.ValidateKeyInput) error {
	log.Println("[ValidateKeyUseCase] Start validate key use case")

	if forwardInput.GivenKey == "" {
		return keycheckerentity.ErrKeyEmpty
	}

	if forwardInput.GivenKey != forwardInput.ValidKey {
		return keycheckerentity.ErrKeyNotValid
	}

	log.Println("[ValidateKeyUseCase] Key is valid")

	return nil
}
