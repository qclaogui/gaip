// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

// Operations provides methods for managing the long-running operations.
// You don't need to initiate this struct. Create a client instance via NewClient, and
// then access Operations through client.Operations field.
type Operations struct {
	apiClient *apiClient
}
