package jwt

import (
	"errors"
	"fmt"
)

var (
	ErrJwtError           = errors.New("jwt error")
	ErrInvalidTokenFormat = fmt.Errorf("%w: invalid token format", ErrJwtError)
	ErrExpiredToken       = fmt.Errorf("%w: expired token", ErrJwtError)
	ErrInvalidSignature   = fmt.Errorf("%w: invalid signature", ErrJwtError)
)
