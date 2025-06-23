package a2a

import (
	"context"
	"errors"
	"io"

	pb "github.com/qclaogui/gaip/genproto/a2a/apiv1/a2apb"
	"google.golang.org/api/iterator"
)

type Tasks struct {
	apiClient *apiClient
}

func (t Tasks) GetTask(ctx context.Context) (*Task, error) {
	req := &pb.GetTaskRequest{}
	debugPrint(req)
	res, err := t.apiClient.a2aClient.GetTask(ctx, req)
	if err != nil {
		return nil, err
	}
	return fromProto[Task](res)
}

func (t Tasks) CancelTask(ctx context.Context) (*Task, error) {
	req := &pb.CancelTaskRequest{}
	debugPrint(req)
	res, err := t.apiClient.a2aClient.CancelTask(ctx, req)
	if err != nil {
		return nil, err
	}
	return fromProto[Task](res)
}

func (t Tasks) TaskSubscription(ctx context.Context) *TaskSubscriptionResponseIterator {
	iter := &TaskSubscriptionResponseIterator{}
	req := &pb.TaskSubscriptionRequest{}
	iter.sc, iter.err = t.apiClient.a2aClient.TaskSubscription(ctx, req)
	return iter
}

// TaskSubscriptionResponseIterator is an iterator over StreamResponse.
type TaskSubscriptionResponseIterator struct {
	sc  pb.A2AService_TaskSubscriptionClient
	err error
}

func (iter *TaskSubscriptionResponseIterator) Next() (*StreamResponse, error) {
	if iter.err != nil {
		return nil, iter.err
	}

	res, err := iter.sc.Recv()
	iter.err = err
	if errors.Is(err, io.EOF) {
		return nil, iterator.Done
	}
	if err != nil {
		return nil, err
	}

	sp, err := fromProto[StreamResponse](res)
	if err != nil {
		iter.err = err
		return nil, err
	}
	return sp, nil
}
