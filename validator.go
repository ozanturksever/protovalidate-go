// Copyright 2023-2024 Buf Technologies, Inc.
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

package protovalidate

import (
	"fmt"
	"sync"

	pvcel "github.com/bufbuild/protovalidate-go/cel"
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var getGlobalValidator = sync.OnceValues(func() (Validator, error) { return New() })

// Validator performs validation on any proto.Message values. The Validator is
// safe for concurrent use.

type Validator interface {
	// Validate checks that message satisfies its constraints. Constraints are
	// defined within the Protobuf file as options from the buf.validate
	// package. An error is returned if the constraints are violated
	// (ValidationError), the evaluation logic for the message cannot be built
	// (CompilationError), or there is a type error when attempting to evaluate
	// a CEL expression associated with the message (RuntimeError).
	Validate(msg proto.Message, options ...ValidationOption) error
}

// New creates a Validator with the given options. An error may occur in setting
// up the CEL execution environment if the configuration is invalid. See the
// individual ValidatorOption for how they impact the fallibility of New.
func New(options ...ValidatorOption) (Validator, error) {
	cfg := config{
		extensionTypeResolver: protoregistry.GlobalTypes,
	}
	for _, opt := range options {
		opt.applyToValidator(&cfg)
	}

	envOpts := []cel.EnvOption{cel.Lib(pvcel.NewLibrary())}
	if cfg.celExtraOptions != nil {
		envOpts = append(envOpts, cfg.celExtraOptions...)
	}
	env, err := cel.NewEnv(envOpts...)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to construct CEL environment: %w", err)
	}

	bldr := newBuilder(
		env,
		cfg.disableLazy,
		cfg.extensionTypeResolver,
		cfg.allowUnknownFields,
		cfg.desc...,
	)

	return &validator{
		failFast: cfg.failFast,
		builder:  bldr,
	}, nil
}

type validator struct {
	builder  *builder
	failFast bool
}

func (v *validator) Validate(
	msg proto.Message,
	options ...ValidationOption,
) error {
	if msg == nil {
		return nil
	}
	cfg := validationConfig{
		failFast: v.failFast,
		filter:   nopFilter{},
	}
	for _, opt := range options {
		opt.applyToValidation(&cfg)
	}
	refl := msg.ProtoReflect()
	eval := v.builder.Load(refl.Descriptor())
	err := eval.EvaluateMessage(refl, &cfg)
	finalizeViolationPaths(err)
	return err
}

// Validate uses a global instance of Validator constructed with no ValidatorOptions and
// calls its Validate function. For the vast majority of validation cases, using this global
// function is safe and acceptable. If you need to provide i.e. a custom
// ExtensionTypeResolver, you'll need to construct a Validator.
func Validate(msg proto.Message) error {
	globalValidator, err := getGlobalValidator()
	if err != nil {
		return err
	}
	return globalValidator.Validate(msg)
}

type config struct {
	failFast              bool
	disableLazy           bool
	desc                  []protoreflect.MessageDescriptor
	extensionTypeResolver protoregistry.ExtensionTypeResolver
	allowUnknownFields    bool
	celExtraOptions       []cel.EnvOption
}

type validationConfig struct {
	failFast bool
	filter   Filter
}
