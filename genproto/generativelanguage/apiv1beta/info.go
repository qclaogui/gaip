// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// SetGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Also passes any
// provided key-value pairs. Intended for use by Google-written clients.
//
// Internal use only.

package generativelanguage

func (c *CacheClient) SetGoogleClientInfo(keyval ...string) {
	c.setGoogleClientInfo(keyval...)
}

func (c *DiscussClient) SetGoogleClientInfo(keyval ...string) {
	c.setGoogleClientInfo(keyval...)
}

func (c *FileClient) SetGoogleClientInfo(keyval ...string) {
	c.setGoogleClientInfo(keyval...)
}

func (c *GenerativeClient) SetGoogleClientInfo(keyval ...string) {
	c.setGoogleClientInfo(keyval...)
}

func (c *ModelClient) SetGoogleClientInfo(keyval ...string) {
	c.setGoogleClientInfo(keyval...)
}

func (c *PermissionClient) SetGoogleClientInfo(keyval ...string) {
	c.setGoogleClientInfo(keyval...)
}

func (c *PredictionClient) SetGoogleClientInfo(keyval ...string) {
	c.setGoogleClientInfo(keyval...)
}

func (c *RetrieverClient) SetGoogleClientInfo(keyval ...string) {
	c.setGoogleClientInfo(keyval...)
}

func (c *TextClient) SetGoogleClientInfo(keyval ...string) {
	c.setGoogleClientInfo(keyval...)
}
