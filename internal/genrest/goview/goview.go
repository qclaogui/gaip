// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package goview

import (
	"fmt"
	"strings"
)

// View contains a list of files to be output.
type View struct {
	Files []*SourceFile
}

// New returns a new, empty View.
func New(capacity int) *View {
	return &View{Files: make([]*SourceFile, 0, capacity)}
}

// Append appends file to this View.
func (view *View) Append(file *SourceFile) *SourceFile {
	view.Files = append(view.Files, file)
	return file
}

// SourceFile contains a single file to be output, including both its content and location.
type SourceFile struct {
	Directory string
	Name      string // without any directory components
	source    *Source
}

// NewFile creates an empty SourceFile with the specified Directory and Name.
func NewFile(directory, name string) *SourceFile {
	return &SourceFile{
		Directory: directory,
		Name:      name,
		source:    NewSource(),
	}
}

// Contents returns the stringifies contents this SourceFile.
func (sf *SourceFile) Contents() string {
	return sf.source.Contents()
}

// Append appends the lines in Source to the lines in SourceFile.
func (sf *SourceFile) Append(source *Source) {
	sf.source.lines = append(sf.source.lines, source.lines...)
}

// P appends a printf-formatted line to SourceFile.
func (sf *SourceFile) P(format string, args ...interface{}) {
	sf.source.P(format, args...)
}

// Source stores a list of lines, typically a continuous section of source code that forms a part of
// (or the entirety of) a whole source file.
type Source struct {
	lines []string // a list of lines
}

// NewSource returns a new Source
func NewSource() *Source {
	return &Source{lines: []string{}}
}

// Contents returns the stringifies contents this SourceFile.
func (source *Source) Contents() string {
	return strings.Join(source.lines, "\n") + "\n"
}

// P writes a new line of content to this SourceFile. The arguments are treated exactly as in
// fmt.Printf. Note that there is an implicit in the SourceFile contents "\n" after each call to
// P().
func (source *Source) P(format string, args ...interface{}) {
	source.lines = append(source.lines, fmt.Sprintf(format, args...))
}
