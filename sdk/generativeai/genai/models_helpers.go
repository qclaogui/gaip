// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

// Text returns a slice of Content with a single Part with the given text.
func Text(text string) []*Content {
	return []*Content{
		NewContentFromText(text, RoleUser),
	}
}

func (c *GenerationConfig) setDefaults() {
	if c == nil {
		return
	}
	// if c.SystemInstruction != nil && c.SystemInstruction.Role == "" {
	// 	c.SystemInstruction.setDefaults()
	// }
}

// func (c *Content) setDefaults() {
// 	if c == nil {
// 		return
// 	}
// 	if c.Role == "" {
// 		c.Role = RoleUser
// 	}
// }
