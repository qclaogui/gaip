// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

// Preview. Live serves as the entry point for establishing real-time WebSocket
// connections to the API. It manages the initial handshake and setup process.
//
// It is initiated when creating a client via [NewClient]. You don't need to
// create a new Live object directly. Access it through the `Live` field of a
// `Client` instance.
//
//	client, _ := genai.NewClient(ctx, &genai.ClientConfig{})
//	session, _ := client.Live.Connect(ctx, model, &genai.LiveConnectConfig{}).
type Live struct {
	apiClient *apiClient
}
