package service

import "errors"

var (
	ErrAlreadyExists = errors.New("item already exists")
	ErrNotExists     = errors.New("item not exists")
	ErrCopyingItem   = errors.New("error copying item")
)
