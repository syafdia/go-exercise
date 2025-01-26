package user

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"time"

	"github.com/syafdia/clean-arch-ddd-cqrs-es/internal/domain"
)

type UserID int64

type User struct {
	ID             UserID
	Email          Email
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewUser(
	email Email,
	password Password,
	createdAt time.Time,
) (*User, error) {
	err := email.Validate()
	if err != nil {
		return nil, err
	}

	err = password.Validate()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := password.Hash()
	if err != nil {
		return nil, err
	}

	return &User{
		Email:          email,
		HashedPassword: hashedPassword,
		CreatedAt:      createdAt,
	}, nil
}

func (u *User) UpdateEmail(newEmail Email, updatedAt time.Time) error {
	err := newEmail.Validate()
	if err != nil {
		return err
	}

	u.Email = newEmail
	u.UpdatedAt = updatedAt

	return nil
}

func (u *User) UpdatePassword(newPassword Password, updatedAt time.Time) error {
	err := newPassword.Validate()
	if err != nil {
		return err
	}

	hashedPassword, err := newPassword.Hash()
	if err != nil {
		return err
	}

	u.HashedPassword = hashedPassword
	u.UpdatedAt = updatedAt

	return nil
}

type Password string

func (p Password) Validate() error {
	if len(p) < 8 {
		return domain.NewValidationError(ErrCodeInvalidPasswordLength, "must be at least 6 characters")
	}

	if len(p) > 16 {
		return domain.NewValidationError(ErrCodeInvalidPasswordLength, "the maximum value is 16 characters")
	}

	return nil
}

func (p Password) Hash() (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(p))
	if err != nil {
		return "", fmt.Errorf("password: failed hashing password, %w", err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

type Email string

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (e Email) Validate() error {
	ok := emailRegex.MatchString(string(e))
	if !ok {
		return domain.NewValidationError(ErrCodeInvalidEmailFormat, "value is not valid")
	}

	return nil
}
