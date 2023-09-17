// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package todo

import (
	"errors"

	"github.com/qclaogui/golang-api-server/genproto/todo/apiv1/todopb"
)

var (
	// ErrNotFound is returned when a item is not found.
	ErrNotFound = errors.New("the item was not found in the repository")

	// ErrFailedToCreate is returned when a item is create Failed
	ErrFailedToCreate = errors.New("failed to add the todo to the repository")
)

// Repository is a interface that defines the rules around what a customer repository
// Has to be able to perform
type Repository interface {
	todopb.ToDoServiceServer
}
