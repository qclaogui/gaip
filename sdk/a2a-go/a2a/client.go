package a2a

import (
	"context"
	"errors"
	"os"
)

// A Client is a A2A client.
type Client struct {
	AgentCards *AgentCards
	Messages   *Messages
	Tasks      *Tasks

	NotificationConfigs *NotificationConfigs
}

// ClientConfig is the configuration for the A2A client.
type ClientConfig struct {
	envVarProvider func() map[string]string
}

func defaultEnvVarProvider() map[string]string {
	vars := make(map[string]string)
	if v, ok := os.LookupEnv("A2A_API_KEY"); ok {
		vars["A2A_API_KEY"] = v
	}
	return vars
}

// NewClient creates a new A2A client with the provided configuration.
func NewClient(ctx context.Context, cc *ClientConfig) (*Client, error) {
	if cc == nil {
		cc = &ClientConfig{}
	}
	if cc.envVarProvider == nil {
		cc.envVarProvider = defaultEnvVarProvider
	}
	envVars := cc.envVarProvider()
	_ = envVars

	ac, err := newAPIClient(ctx)
	if err != nil {
		return nil, err
	}

	c := &Client{
		AgentCards:          &AgentCards{apiClient: ac},
		Messages:            &Messages{apiClient: ac},
		NotificationConfigs: &NotificationConfigs{apiClient: ac},
	}
	return c, nil
}

func (c *Client) Close() error {
	return errors.Join(
		c.AgentCards.apiClient.a2aClient.Close(),
		c.Messages.apiClient.a2aClient.Close(),
		c.NotificationConfigs.apiClient.a2aClient.Close(),
	)
}

func fromProto[V interface{ fromProto(P) *V }, P any](p P) (*V, error) {
	var v V
	return pvCatchPanic(func() *V { return v.fromProto(p) })
}
