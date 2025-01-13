// Copyright 2024, KhulnaSoft Ltd.
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
//go:build !all
// +build !all

package main

import (
	"fmt"

	"github.com/khulnasoft/codeinfra/pkg/v3/resource/provider"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	perrors "github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/errors"
	codeinfraprovider "github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/provider"
)

type Component struct {
	codeinfra.ResourceState

	Foo codeinfra.StringOutput `codeinfra:"foo"`
}

type ComponentArgs struct {
	Foo codeinfra.StringInput `codeinfra:"foo"`
}

func NewComponent(ctx *codeinfra.Context, name string, args *ComponentArgs, opts ...codeinfra.ResourceOption) (*Component, error) {
	component := &Component{}
	err := ctx.RegisterComponentResource("testcomponent:index:Component", name, component, opts...)
	if err != nil {
		return nil, err
	}

	if err := ctx.RegisterResourceOutputs(component, codeinfra.Map{
		"foo": args.Foo,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

const (
	providerName = "testcomponent"
	version      = "0.0.1"
)

func main() {
	if err := provider.MainWithOptions(provider.Options{
		Name:    providerName,
		Version: version,
		Construct: func(ctx *codeinfra.Context, typ, name string, inputs codeinfraprovider.ConstructInputs,
			options codeinfra.ResourceOption,
		) (*codeinfraprovider.ConstructResult, error) {
			if typ != "testcomponent:index:Component" {
				return nil, fmt.Errorf("unknown resource type %s", typ)
			}

			args := &ComponentArgs{}
			if err := inputs.CopyTo(args); err != nil {
				return nil, fmt.Errorf("setting args: %w", err)
			}

			component, err := NewComponent(ctx, name, args, options)
			if err != nil {
				return nil, fmt.Errorf("creating component: %w", err)
			}

			return codeinfraprovider.NewConstructResult(component)
		},
		Call: func(ctx *codeinfra.Context, tok string, args codeinfraprovider.CallArgs) (*codeinfraprovider.CallResult, error) {
			if tok != "testcomponent:index:Component/getMessage" {
				return nil, fmt.Errorf("unknown method %s", tok)
			}

			return nil, perrors.NewInputPropertyError("foo", "the failure reason")
		},
	}); err != nil {
		cmdutil.ExitError(err.Error())
	}
}
