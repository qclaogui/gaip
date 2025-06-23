// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package a2a

import (
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
)

// WithREST is an option that enables REST transport for the client.
// The default transport (if this option isn't provided) is gRPC.
func WithREST() option.ClientOption {
	return &withREST{}
}

type withREST struct {
	internaloption.EmbeddableAdapter
}

func (w *withREST) applyOpt(c *config) {
	c.withREST = true
}

// WithClientInfo is an option that sets request information
// identifying the product that is calling this client.
func WithClientInfo(key, value string) option.ClientOption {
	return &withClientInfo{key: key, value: value}
}

type withClientInfo struct {
	internaloption.EmbeddableAdapter
	key, value string
}

func (w *withClientInfo) applyOpt(c *config) {
	c.ciKey = w.key
	c.ciValue = w.value
}

type config struct {
	// withREST tells the client to use REST as the underlying transport.
	withREST bool
	// key-value pair to add to the Google client info header.
	ciKey   string
	ciValue string
}

// newConfig generates a new config with all the given
// clientOption applied.
func newConfig(opts ...option.ClientOption) config {
	var conf config
	for _, opt := range opts {
		if vOpt, ok := opt.(clientOption); ok {
			vOpt.applyOpt(&conf)
		}
	}
	return conf
}

// A clientOption is an option for a client.
type clientOption interface {
	option.ClientOption
	applyOpt(*config)
}
