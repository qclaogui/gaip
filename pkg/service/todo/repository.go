package todo

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	// ErrTodoNotFound is returned when a item is not found.
	ErrTodoNotFound = errors.New("the todo item was not found in the repository")

	// ErrTodoFailedToCreate is returned when a item is create Failed
	ErrTodoFailedToCreate = errors.New("failed to add the todo to the repository")
)

// Repository is a interface that defines the rules around what a customer repository
// Has to be able to perform
type Repository interface {
	Read(context.Context, uuid.UUID) (*Todo, error)

	Create(context.Context, *Todo) (*Todo, error)

	Update(context.Context, *Todo) (int64, error)

	Delete(context.Context, uuid.UUID) (int64, error)

	ReadAll(context.Context) ([]*Todo, error)
}
