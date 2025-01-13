// Copyright 2016-2021, KhulnaSoft Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codegentest

import (
	"context"
	"fmt"
	"plain-and-default/foo"
	"testing"
	"time"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/resource"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/stretchr/testify/assert"
)

type mocks int

// Create the mock.
func (mocks) NewResource(args codeinfra.MockResourceArgs) (string, resource.PropertyMap, error) {
	return args.Name, args.Inputs, nil
}

func (mocks) Call(args codeinfra.MockCallArgs) (resource.PropertyMap, error) {
	panic("functions not supported")
}

func TestDefaults(t *testing.T) {
	codeinfraTest(t, "explicit false", func(ctx *codeinfra.Context) error {
		output, err := foo.NewModuleResource(ctx, "test", &foo.ModuleResourceArgs{
			OptionalBool: codeinfra.Bool(false),
		})
		assert.NoError(t, err)
		assert.Equalf(t, *waitOut(t, output.OptionalBool).(*bool), false,
			"Value has been set to false, make sure it doesn't change.")
		return nil
	})

	codeinfraTest(t, "explicit true", func(ctx *codeinfra.Context) error {
		output, err := foo.NewModuleResource(ctx, "test", &foo.ModuleResourceArgs{
			OptionalBool: codeinfra.Bool(true),
		})
		assert.NoError(t, err)
		assert.Equalf(t, *waitOut(t, output.OptionalBool).(*bool), true,
			"Value has been set to true, make sure it doesn't change.")
		return nil
	})

	codeinfraTest(t, "default value", func(ctx *codeinfra.Context) error {
		output, err := foo.NewModuleResource(ctx, "test", &foo.ModuleResourceArgs{})
		assert.NoError(t, err)
		assert.Equalf(t, *waitOut(t, output.OptionalBool).(*bool), true,
			"Default value is true, and the value has not been specified")
		return nil
	})
}

func codeinfraTest(t *testing.T, name string, testBody func(*codeinfra.Context) error) {
	t.Run(name, func(t *testing.T) {
		err := codeinfra.RunErr(testBody, codeinfra.WithMocks("project", "stack", mocks(0)))
		assert.NoError(t, err)
	})
}

func waitOut(t *testing.T, output codeinfra.Output) interface{} {
	result, err := waitOutput(output, 1*time.Second)
	if !assert.NoError(t, err, "output not received") {
		return nil
	}
	return result
}

func waitOutput(output codeinfra.Output, timeout time.Duration) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := make(chan interface{})
	output.ApplyT(func(v interface{}) interface{} {
		ch <- v
		return v
	})

	select {
	case v := <-ch:
		return v, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("timed out waiting for codeinfra.Output after %v: %w", timeout, ctx.Err())
	}
}
