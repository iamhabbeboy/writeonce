package entity

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

var ErrPipeNotFound = NewEntityError("Pipe", ErrNotFound)
var ErrProjectNotFound = NewEntityError("Project", ErrNotFound)
var ErrTemplateNotFound = NewEntityError("Template", ErrNotFound)

type EntityError struct {
	Name string
	Err  error
}

func (e *EntityError) Unwrap() error {
	return e.Err
}

func (e *EntityError) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Err)
}

func NewEntityError(name string, err error) *EntityError {
	return &EntityError{
		Name: name,
		Err:  err,
	}
}
