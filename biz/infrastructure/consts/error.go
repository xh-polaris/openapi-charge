package consts

import "errors"

var (
	ErrInValidId          = errors.New("invalid objectId")
	ErrNotFound           = errors.New("object not found")
	ErrNoBaseInf          = errors.New("no Base Interface found")
	ErrNoFullInf          = errors.New("no Full Interface found")
	ErrRepetition         = errors.New("create repetition")
	ErrUpdate             = errors.New("update failed")
	ErrInsufficientMargin = errors.New("insufficient margin")
	ErrAccount            = errors.New("insert account error")
)
