// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/option"

	aiplatform "github.com/qclaogui/gaip/genproto/aiplatform/apiv1beta1"
	"github.com/qclaogui/gaip/vertexai/internal"
)

// A Client is a Google Vertex AI client.
type Client struct {
	pc *aiplatform.PredictionClient
	cc *cacheClient

	projectID string
	location  string
}

// NewClient creates a new Google Vertex AI client.
//
// Clients should be reused instead of created as needed. The methods of Client
// are safe for concurrent use by multiple goroutines.
// projectID is your GCP project; location is GCP region/location per
// https://cloud.google.com/vertex-ai/docs/general/locations
// If location is empty, this function attempts to infer it from environment
// variables and falls back to a default location if unsuccessful.
//
// You may configure the client by passing in options from the
// [google.golang.org/api/option] package. You may also use options defined in
// this package, such as [WithREST].
func NewClient(ctx context.Context, projectID, location string, opts ...option.ClientOption) (*Client, error) {
	location = inferLocation(location)
	endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)
	conf := newConfig(opts...)
	return newClient(ctx, projectID, location, endpoint, conf, opts)
}

func newClient(ctx context.Context, projectID, location, endpoint string, conf config, opts []option.ClientOption) (*Client, error) {
	opts = append([]option.ClientOption{option.WithEndpoint(endpoint)}, opts...)
	c := &Client{projectID: projectID, location: location}

	if err := setGAPICClient(ctx,
		&c.pc,
		conf,
		aiplatform.NewPredictionRESTClient,
		aiplatform.NewPredictionClient, opts); err != nil {
		return nil, err
	}

	if err := setGAPICClient(ctx,
		&c.cc,
		conf,
		newCacheRESTClient,
		newCacheClient, opts); err != nil {
		return nil, err
	}

	return c, nil
}

type sgci interface{ SetGoogleClientInfo(...string) }

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
	c.SetGoogleClientInfo(kvs...)
	*pf = c
	return nil
}

// Close closes the client.
func (c *Client) Close() error {
	pcErr := c.pc.Close()
	ccErr := c.cc.Close()

	switch {
	case pcErr != nil:
		return pcErr
	case ccErr != nil:
		return ccErr
	default:
		return nil
	}
}

const defaultLocation = "us-central1"

// inferLocation infers the GCP location from its parameter, env vars or
// a default location.
func inferLocation(location string) string {
	if location != "" {
		return location
	}
	if location = os.Getenv("GOOGLE_CLOUD_REGION"); location != "" {
		return location
	}
	if location = os.Getenv("CLOUD_ML_REGION"); location != "" {
		return location
	}

	return defaultLocation
}

func int32pToFloat32p(x *int32) *float32 {
	if x == nil {
		return nil
	}
	f := float32(*x)
	return &f
}

func float32pToInt32p(x *float32) *int32 {
	if x == nil {
		return nil
	}
	i := int32(*x)
	return &i
}
