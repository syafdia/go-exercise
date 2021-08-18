package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Create(ctx context.Context, input CreateUserInput) (User, error)
	Update(ctx context.Context, input UpdateUserInput) (User, error)
	FindOneByID(ctx context.Context, id int64) (User, error)
	Destroy(ctx context.Context, id int64) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

func (u *userRepo) Create(ctx context.Context, input CreateUserInput) (User, error) {
	result, err := u.db.ExecContext(ctx,
		`INSERT INTO users (email, first_name, last_name, gender) VALUES ($1, $2, $3, $4);`,
		input.Email,
		input.FirstName,
		input.LastName,
		input.Gender)
	if err != nil {
		return User{}, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}

	return User{
		ID:        userID,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Gender:    input.Gender,
	}, nil
}

func (u *userRepo) Update(ctx context.Context, input UpdateUserInput) (User, error) {
	// TODO
	return User{}, nil
}

func (u *userRepo) FindOneByID(ctx context.Context, id int64) (User, error) {
	return User{}, nil
}

func (u *userRepo) Destroy(ctx context.Context, id int64) error {
	// TODO
	return nil
}
