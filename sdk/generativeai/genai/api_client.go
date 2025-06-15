// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import (
	"context"

	gl "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta"
	"github.com/qclaogui/gaip/sdk/generativeai/internal"
	"google.golang.org/api/option"
)

type apiClient struct {
	// clientConfig *ClientConfig

	gc *gl.GenerativeClient
	cc *gl.CacheClient
	fc *gl.FileClient
	mc *gl.ModelClient
}

func newAPIClient(ctx context.Context, opts ...option.ClientOption) (*apiClient, error) {
	ac := &apiClient{}
	conf := newConfig(opts...)

	if err := setGAPICClient(ctx, &ac.gc, conf, gl.NewGenerativeRESTClient, gl.NewGenerativeClient, opts); err != nil {
		return nil, err
	}

	if err := setGAPICClient(ctx, &ac.cc, conf, gl.NewCacheRESTClient, gl.NewCacheClient, opts); err != nil {
		return nil, err
	}

	if err := setGAPICClient(ctx, &ac.fc, conf, gl.NewFileRESTClient, gl.NewFileClient, opts); err != nil {
		return nil, err
	}

	if err := setGAPICClient(ctx, &ac.mc, conf, gl.NewModelRESTClient, gl.NewModelClient, opts); err != nil {
		return nil, err
	}

	return ac, nil
}

type sgci interface {
	SetGoogleClientInfo(...string)
}

func setGAPICClient[ClientType sgci](ctx context.Context, pf *ClientType, conf config, newREST, newGRPC func(context.Context, ...option.ClientOption) (ClientType, error), opts []option.ClientOption) error {
	var c ClientType
	var err error
	if conf.withREST {
		c, err = newREST(ctx, opts...)
	} else {
		c, err = newGRPC(ctx, opts...)
	}
	if err != nil {
		return err
	}

	kvs := []string{"gccl", internal.Version}
	if conf.ciKey != "" && conf.ciValue != "" {
		kvs = append(kvs, conf.ciKey, conf.ciValue)
	}
	_ = kvs

	c.SetGoogleClientInfo(kvs...)
	*pf = c
	return nil
}

// Close closes the client.
func (c *apiClient) Close() error {
	pcErr := c.gc.Close()
	ccErr := c.cc.Close()
	fcErr := c.fc.Close()
	mcErr := c.mc.Close()

	switch {
	case pcErr != nil:
		return pcErr
	case ccErr != nil:
		return ccErr
	case fcErr != nil:
		return fcErr
	case mcErr != nil:
		return mcErr
	default:
		return nil
	}
}
