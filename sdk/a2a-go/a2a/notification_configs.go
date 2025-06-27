// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package a2a

import (
	"context"

	a2a "github.com/qclaogui/gaip/genproto/a2a/apiv1"
	pb "github.com/qclaogui/gaip/genproto/a2a/apiv1/a2apb"
	"google.golang.org/api/iterator"
)

type NotificationConfigs struct {
	apiClient *apiClient
}

func (n NotificationConfigs) CreateNotificationConfig(ctx context.Context) (*TaskPushNotificationConfig, error) {
	req := &pb.CreateTaskPushNotificationConfigRequest{}
	debugPrint(req)
	res, err := n.apiClient.a2aClient.CreateTaskPushNotificationConfig(ctx, req)
	if err != nil {
		return nil, err
	}
	return fromProto[TaskPushNotificationConfig](res)
}

func (n NotificationConfigs) GetNotificationConfig(ctx context.Context) (*TaskPushNotificationConfig, error) {
	req := &pb.GetTaskPushNotificationConfigRequest{}
	debugPrint(req)
	res, err := n.apiClient.a2aClient.GetTaskPushNotificationConfig(ctx, req)
	if err != nil {
		return nil, err
	}
	return fromProto[TaskPushNotificationConfig](res)
}

// DeleteNotificationConfig deletes the NotificationConfig with the given name.
func (n NotificationConfigs) DeleteNotificationConfig(ctx context.Context, name string) error {
	return n.apiClient.a2aClient.DeleteTaskPushNotificationConfig(ctx, &pb.DeleteTaskPushNotificationConfigRequest{Name: name})
}

func (n NotificationConfigs) ListNotificationConfigs(ctx context.Context) *NotificationConfigIterator {
	return &NotificationConfigIterator{
		iter: n.apiClient.a2aClient.ListTaskPushNotificationConfig(ctx, &pb.ListTaskPushNotificationConfigRequest{}),
	}
}

type NotificationConfigIterator struct {
	iter *a2a.TaskPushNotificationConfigIterator
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NotificationConfigIterator) Next() (*TaskPushNotificationConfig, error) {
	nc, err := it.iter.Next()
	if err != nil {
		return nil, err
	}
	return fromProto[TaskPushNotificationConfig](nc)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NotificationConfigIterator) PageInfo() *iterator.PageInfo {
	return it.iter.PageInfo()
}

func (p *PushNotificationConfig) setDefaults() {
	if p == nil {
		return
	}
	if p.Authentication != nil {
		p.Authentication.setDefaults()
	}
}

func (a *AuthenticationInfo) setDefaults() {
	if a == nil {
		return
	}
}
