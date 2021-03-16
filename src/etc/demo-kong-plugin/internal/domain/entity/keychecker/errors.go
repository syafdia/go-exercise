package keychecker

import (
	"errors"
)

var (
	ErrKeyEmpty    = errors.New("key is empty")
	ErrKeyNotValid = errors.New("key is not valid")
)
