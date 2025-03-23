package auth

import "errors"

var (
	ErrNotFound = errors.New("data is not found")
)

type User struct {
	ID       string
	Username string
	Password string
}
