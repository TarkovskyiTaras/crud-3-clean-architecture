package entity

import "errors"

var ErrNotFound = errors.New("not found")
var ErrInvalidEntity = errors.New("invalid entity")
var ErrConflict = errors.New("item already exists")
