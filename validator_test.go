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
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"google.golang.org/protobuf/reflect/protoregistry"
	"testing"

	pb "github.com/bufbuild/protovalidate-go/internal/gen/tests/example/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/apipb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/sourcecontextpb"
)

func TestValidator_Validate(t *testing.T) {
	t.Parallel()

	t.Run("HasMsgExprs", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)

		tests := []struct {
			msg   *pb.HasMsgExprs
			exErr bool
		}{
			{
				&pb.HasMsgExprs{X: 2, Y: 43},
				false,
			},
			{
				&pb.HasMsgExprs{X: 9, Y: 8},
				true,
			},
		}

		for _, test := range tests {
			err := val.Validate(test.msg)
			if test.exErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		}
	})
}

func TestRecursive(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)

	selfRec := &pb.SelfRecursive{X: 123, Turtle: &pb.SelfRecursive{X: 456}}
	err = val.Validate(selfRec)
	require.NoError(t, err)

	loopRec := &pb.LoopRecursiveA{B: &pb.LoopRecursiveB{}}
	err = val.Validate(loopRec)
	require.NoError(t, err)
}

func TestValidator_ValidateOneof(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	oneofMessage := &pb.MsgHasOneof{O: &pb.MsgHasOneof_X{X: "foo"}}
	err = val.Validate(oneofMessage)
	require.NoError(t, err)

	oneofMessage = &pb.MsgHasOneof{O: &pb.MsgHasOneof_Y{Y: 42}}
	err = val.Validate(oneofMessage)
	require.NoError(t, err)

	oneofMessage = &pb.MsgHasOneof{O: &pb.MsgHasOneof_Msg{Msg: &pb.HasMsgExprs{X: 4, Y: 50}}}
	err = val.Validate(oneofMessage)
	require.NoError(t, err)

	oneofMessage = &pb.MsgHasOneof{}
	err = val.Validate(oneofMessage)
	assert.Error(t, err)
}

func TestValidator_ValidateRepeatedFoo(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	repeatMessage := &pb.MsgHasRepeated{
		X: []float32{1, 2, 3},
		Y: []string{"foo", "bar"},
		Z: []*pb.HasMsgExprs{
			{
				X: 4,
				Y: 55,
			}, {
				X: 4,
				Y: 60,
			},
		},
	}
	err = val.Validate(repeatMessage)
	require.NoError(t, err)
}

func TestValidator_ValidateMapFoo(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	mapMessage := &pb.MsgHasMap{
		Int32Map:   map[int32]int32{-1: 1, 2: 2},
		StringMap:  map[string]string{"foo": "foo", "bar": "bar", "baz": "baz"},
		MessageMap: map[int64]*pb.LoopRecursiveA{0: nil},
	}
	err = val.Validate(mapMessage)
	require.Error(t, err)
}

func TestValidator_Validate_TransitiveFieldConstraints(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := &pb.TransitiveFieldConstraint{
		Mask: &fieldmaskpb.FieldMask{Paths: []string{"foo", "bar"}},
	}
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_MultipleStepsTransitiveFieldConstraints(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := &pb.MultipleStepsTransitiveFieldConstraints{
		Api: &apipb.Api{
			SourceContext: &sourcecontextpb.SourceContext{
				FileName: "path/file",
			},
		},
	}
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_FieldOfTypeAny(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	simple := &pb.Simple{S: "foo"}
	anyFromSimple, err := anypb.New(simple)
	require.NoError(t, err)
	msg := &pb.FieldOfTypeAny{
		Any: anyFromSimple,
	}
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_CelMapOnARepeated(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := &pb.CelMapOnARepeated{Values: []*pb.CelMapOnARepeated_Value{
		{Name: "foo"},
		{Name: "bar"},
		{Name: "baz"},
	}}
	err = val.Validate(msg)
	require.NoError(t, err)
	msg.Values = append(msg.Values, &pb.CelMapOnARepeated_Value{Name: "foo"})
	err = val.Validate(msg)
	valErr := &ValidationError{}
	require.ErrorAs(t, err, &valErr)
}

func TestValidator_Validate_RepeatedItemCel(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := &pb.RepeatedItemCel{Paths: []string{"foo"}}
	err = val.Validate(msg)
	require.NoError(t, err)
	msg.Paths = append(msg.Paths, " bar")
	err = val.Validate(msg)
	valErr := &ValidationError{}
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "paths.no_space", valErr.Violations[0].GetConstraintId())
}

func TestValidator_Validate_CustomCelEnv(t *testing.T) {
	t.Parallel()
	e, err := cel.NewEnv(
		cel.TypeDescs(protoregistry.GlobalFiles),
		cel.Lib(DefaultCelLib(true)),
		cel.Function("isCustomCheck",
			cel.Overload(
				"string_is_custom_check_bool",
				[]*cel.Type{cel.StringType},
				cel.BoolType,
				cel.FunctionBinding(func(args ...ref.Val) ref.Val {
					v, ok := args[0].Value().(string)
					if !ok {
						return types.Bool(false)
					}
					if v == "true" {
						return types.Bool(true)
					}
					return types.Bool(false)
				}),
			),
		),
	)
	val, err := New(WithCelEnv(e))
	require.NoError(t, err)
	msg := &pb.CustomCelEnv{Value: "true"}
	err = val.Validate(msg)
	require.NoError(t, err)
	msg.Value = "false"
	err = val.Validate(msg)
	valErr := &ValidationError{}
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "custom", valErr.Violations[0].GetConstraintId())
}
