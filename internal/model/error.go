package model

import (
	"errors"
	"fmt"
)

var (
	ErrPermissionDenied = errors.New("permission denied")
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExist     = errors.New("already exist")
)

var (
	ErrTagNotFound = fmt.Errorf("tag not found: %w", ErrNotFound)
)
