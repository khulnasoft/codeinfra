// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/internals"
)

type componentArgs struct {
	First  string `codeinfra:"first"`
	Second string `codeinfra:"second"`
}

type ComponentArgs struct {
	First  codeinfra.StringInput
	Second codeinfra.StringInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type Component struct {
	codeinfra.ResourceState
}

func NewComponent(
	ctx *codeinfra.Context, name string, args *ComponentArgs, opts ...codeinfra.ResourceOption,
) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

func (c *Component) GetMessage(ctx *codeinfra.Context, args *ComponentGetMessageArgs) (ComponentGetMessageResultOutput, error) {
	out, err := ctx.Call("testcomponent:index:Component/getMessage", args, ComponentGetMessageResultOutput{}, c)
	if err != nil {
		return ComponentGetMessageResultOutput{}, err
	}
	return out.(ComponentGetMessageResultOutput), nil
}

type componentGetMessageArgs struct {
	Name string `codeinfra:"name"`
}

type ComponentGetMessageArgs struct {
	Name codeinfra.StringInput
}

func (ComponentGetMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentGetMessageArgs)(nil)).Elem()
}

type ComponentGetMessageResult struct {
	Message string `codeinfra:"message"`
}

type ComponentGetMessageResultOutput struct{ *codeinfra.OutputState }

func (ComponentGetMessageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentGetMessageResult)(nil)).Elem()
}

func (o ComponentGetMessageResultOutput) Message() codeinfra.StringOutput {
	return o.ApplyT(func(v ComponentGetMessageResult) string { return v.Message }).(codeinfra.StringOutput)
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((*Component)(nil))
}

func init() {
	codeinfra.RegisterOutputType(ComponentGetMessageResultOutput{})
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		component, err := NewComponent(ctx, "component", &ComponentArgs{
			First:  codeinfra.String("Hello"),
			Second: codeinfra.String("World"),
		})
		if err != nil {
			return err
		}
		result, err := component.GetMessage(ctx, &ComponentGetMessageArgs{
			Name: codeinfra.String("Alice"),
		})
		if err != nil {
			return err
		}
		message := result.Message()
		ctx.Export("message", message)
		ctx.Export("messagedeps", awaitDependencies(ctx, message))

		return nil
	})
}

func awaitDependencies(ctx *codeinfra.Context, o codeinfra.Output) codeinfra.URNArray {
	r, err := internals.UnsafeAwaitOutput(ctx.Context(), o)
	if err != nil {
		panic(err)
	}
	var deps codeinfra.URNArray
	for _, dep := range r.Dependencies {
		deps = append(deps, dep.URN())
	}
	return deps
}
