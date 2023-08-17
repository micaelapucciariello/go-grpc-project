package service

import "errors"

var (
	ErrAlreadyExists = errors.New("item already exists")
	ErrCopyingItem   = errors.New("error copying item")
)
