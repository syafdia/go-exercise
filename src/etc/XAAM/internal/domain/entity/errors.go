package entity

import (
	"errors"
	"net/http"
)

type Err struct {
	Errr   error
	Status int
}

func NewErr(status int, message string) error {
	return &Err{Status: status, Errr: errors.New(message)}
}

func WrapErr(status int, err error) error {
	return &Err{Status: status, Errr: err}
}

func (e *Err) Error() string {
	return e.Errr.Error()
}

func (e *Err) Unwrap() error {
	return e.Errr
}

var (
	ErrNotFound   = NewErr(http.StatusNotFound, "not found")
	ErrBadRequest = NewErr(http.StatusBadRequest, "bad request")
)
