package domain

import "errors"

var (
	ErrAccountNotFound    = errors.New("account not found")
	ErrDuplicatedApiKey   = errors.New("api key already exists")
	ErrInvalidAccountID   = errors.New("invalid account ID")
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)
