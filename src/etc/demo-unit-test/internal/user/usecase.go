package user

import "context"

type RegisterUserUseCase interface {
	Execute(ctx context.Context, input CreateUserInput) (User, error)
}

type registerUserUseCase struct {
	userRepo UserRepo
}

func NewRegisterUserUseCase(
	userRepo UserRepo,
) RegisterUserUseCase {
	return &registerUserUseCase{userRepo: userRepo}
}

func (w *registerUserUseCase) Execute(ctx context.Context, input CreateUserInput) (User, error) {
	user, err := w.userRepo.Create(ctx, input)
	if err != nil {
		return User{}, err
	}

	// Do something
	// ...
	// ...

	return user, nil
}
