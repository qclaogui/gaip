// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package errorhandling

import (
	"fmt"
	"strings"
)

// Accumulator allows storing a series of errors and then concatenating their string representations
// into one big error.
type Accumulator struct {
	errors []error
}

// AccumulateError stores an error to be reported later.
func (ea *Accumulator) AccumulateError(err error) {
	if err == nil {
		return
	}
	ea.errors = append(ea.errors, err)
}

// Error concatenates the string representations of all stored errors and returns it as a single
// error.
func (ea *Accumulator) Error() error {
	if len(ea.errors) == 0 {
		return nil
	}
	errorStrings := make([]string, len(ea.errors))
	for idx, err := range ea.errors {
		errorStrings[idx] = err.Error()
	}
	return fmt.Errorf(strings.Join(errorStrings, "\n"))
}
