package consts

import "errors"

var (
	ErrInValidId = errors.New("invalid objectId")
	ErrNotFound  = errors.New("object not found")
)
