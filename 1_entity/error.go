package entity

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidEntity = errors.New("invalid entity")
var ErrConflict = errors.New("item already exists")
