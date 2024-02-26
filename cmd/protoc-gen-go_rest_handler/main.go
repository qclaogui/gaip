// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package main

import (
	"flag"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var serviceConfigPath *string

func main() {
	var flags flag.FlagSet
	serviceConfigPath = flags.String("api-service-config", "", "API service config file")

	opts := &protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(
		func(gen *protogen.Plugin) error {
			for _, f := range gen.Files {
				if !f.Generate {
					continue
				}

				if err := GenerateFile(gen, f); err != nil {
					return err
				}
			}
			gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
			return nil
		})
}

var generatedFiles = map[string]bool{}

// GenerateFile generates the contents of a _handler.pb.go file.
func GenerateFile(plugin *protogen.Plugin, file *protogen.File) error {
	if len(file.Services) == 0 {
		return nil
	}

	prefix := filepath.Dir(filepath.ToSlash(file.GeneratedFilenamePrefix))
	if generatedFiles[prefix] {
		return nil
	}
	generatedFiles[prefix] = true

	protoModel, err := NewProtoModel(plugin)
	if err != nil {
		return err
	}

	g := plugin.NewGeneratedFile(
		filepath.Join(prefix, "genrest", "rest_handler_response.txt"),
		file.GoImportPath,
	)

	g.P("Generated via \"google.golang.org/protobuf/compiler/protogen\" via ProtoModel!")
	g.P("\n")
	g.P("Proto Files:\n", strings.Join(plugin.Request.GetFileToGenerate(), "\n"))

	g.P("\nProto Model:")
	g.P(protoModel.String())

	g.P("\n\n")

	goModel, err := NewGoModel(protoModel)
	if err != nil {
		return err
	}
	g.P(goModel.String())

	view, err := NewView(goModel)
	if err != nil {
		return err
	}

	for _, source := range view.Files {
		f := plugin.NewGeneratedFile(
			filepath.Join(prefix, "genrest", source.Name),
			file.GoImportPath,
		)

		f.P(source.Contents())
	}

	return nil
}
