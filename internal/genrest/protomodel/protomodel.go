// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package protomodel

import (
	"fmt"
	"strings"

	"github.com/qclaogui/gaip/internal/genrest/errorhandling"
	"github.com/qclaogui/gaip/internal/genrest/pbinfo"
	"google.golang.org/protobuf/types/descriptorpb"
)

////////////////////////////////////////
// ProtoModel

// Model is a data model encapsulating the relevant information for a REST-proto transcoding for
// various services defined via annotated protocol buffers.
type Model struct {
	errorhandling.Accumulator
	ProtoInfo pbinfo.Info
	Services  []*Service
}

// String returns a string representation of this Model.
func (model *Model) String() string {
	services := make([]string, len(model.Services))
	for idx, svc := range model.Services {
		if svc == nil {
			continue
		}
		services[idx] = svc.String()
	}
	return strings.Join(services, "\n\n")
}

// AddService adds `service` to this Service.
func (model *Model) AddService(service *Service) *Service {
	model.Services = append(model.Services, service)
	return service
}

////////////////////////////////////////
// Service

// Service is a data model encapsulating the information relevant to REST-proto transcoding about a
// proto-defined service.
type Service struct {
	Descriptor   *descriptorpb.ServiceDescriptorProto // maybe not needed
	Name         string
	TypeName     string
	RESTBindings []*RESTBinding
}

// String returns a string representation of this Service.
func (service *Service) String() string {
	handlers := make([]string, len(service.RESTBindings))
	for idx, h := range service.RESTBindings {
		handlers[idx] = h.String()
	}
	indent := "  "
	return fmt.Sprintf("%s (%s):\n%s%s", service.Name, service.TypeName, indent, strings.Join(handlers, "\n"+indent))
}

// AddBinding adds a RESTBinding to this Service.
func (service *Service) AddBinding(binding *RESTBinding) {
	service.RESTBindings = append(service.RESTBindings, binding)
}

////////////////////////////////////////
// RESTBinding

// RESTBinding encapsulates the information contained in a protocol buffer HTTP annotation.
type RESTBinding struct {
	// Index of the binding for this method. Since methods could contain multiple bindings, we
	// will need a way to identify each binding uniquely.
	Index int

	// The name of the method for which this is a binding.
	ProtoMethod string

	// The URL pattern of the binding.
	RESTPattern *RESTRequestPattern

	// The fields in the request body: either none (empty string), a single field (top-level
	// request field, non-dotted), or all not captured in the URL ("*").
	BodyField string
}

// String returns a string representation of this RESTBinding.
func (binding *RESTBinding) String() string {
	return fmt.Sprintf("%s[%d] : %s", binding.ProtoMethod, binding.Index, binding.RESTPattern)
}

////////////////////////////////////////
// RESTRequestPattern

// RESTRequestPattern encapsulates the information in an individual REST binding within an HTTP annotation.
type RESTRequestPattern struct {
	HTTPMethod string // HTTP verb
	Pattern    string // the URL pattern
}

// String returns a string representation of this RESTRequestPattern.
func (binding *RESTRequestPattern) String() string {
	return fmt.Sprintf("%s: %q", binding.HTTPMethod, binding.Pattern)
}
