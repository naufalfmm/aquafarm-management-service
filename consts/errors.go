package consts

import "errors"

var (
	ErrPathNotFound   = errors.New("path not found")
	ErrUnclaimedToken = errors.New("unclaimed token")
	ErrInvalidToken   = errors.New("invalid token")
	ErrIdRequired     = errors.New("id is required")
	ErrRecordNotFound = errors.New("record not found")
)
