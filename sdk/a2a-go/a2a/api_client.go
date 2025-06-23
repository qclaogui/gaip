package a2a

import (
	"context"

	"google.golang.org/api/option"

	a2a "github.com/qclaogui/gaip/genproto/a2a/apiv1"
	"github.com/qclaogui/gaip/sdk/a2a-go/internal"
)

// apiClient is a client for the A2A API.
type apiClient struct {
	a2aClient *a2a.Client
}

func newAPIClient(ctx context.Context, opts ...option.ClientOption) (*apiClient, error) {
	ac := &apiClient{}
	conf := newConfig(opts...)

	if err := setGAPICClient(ctx, &ac.a2aClient, conf, a2a.NewRESTClient, a2a.NewClient, opts); err != nil {
		return nil, err
	}

	return ac, nil
}

type sgci interface {
	SetGoogleClientInfo(...string)
}

// setGAPICClient initializes a Google API Client with the provided configuration.
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
