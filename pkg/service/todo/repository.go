package todo

import (
	"context"
	"errors"

	pb "github.com/qclaogui/golang-api-server/pkg/api/todopb/v1"
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
	Create(context.Context, *pb.CreateRequest) (*pb.CreateResponse, error)

	Read(context.Context, *pb.ReadRequest) (*pb.ReadResponse, error)

	Update(context.Context, *pb.UpdateRequest) (*pb.UpdateResponse, error)

	Delete(context.Context, *pb.DeleteRequest) (*pb.DeleteResponse, error)

	ReadAll(context.Context, *pb.ReadAllRequest) (*pb.ReadAllResponse, error)
}
