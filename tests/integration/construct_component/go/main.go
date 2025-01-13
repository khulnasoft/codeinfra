// Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type componentArgs struct {
	Echo interface{} `codeinfra:"echo"`
}

type ComponentArgs struct {
	Echo codeinfra.Input
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type Component struct {
	codeinfra.ResourceState

	Echo    codeinfra.AnyOutput    `codeinfra:"echo"`
	ChildID codeinfra.StringOutput `codeinfra:"childId"`
	Secret  codeinfra.StringOutput `codeinfra:"secret"`
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

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		componentA, err := NewComponent(ctx, "a", &ComponentArgs{Echo: codeinfra.Int(42)})
		if err != nil {
			return err
		}
		_, err = NewComponent(ctx, "b", &ComponentArgs{Echo: componentA.Echo})
		if err != nil {
			return err
		}
		_, err = NewComponent(ctx, "C", &ComponentArgs{Echo: componentA.ChildID})
		if err != nil {
			return err
		}
		return nil
	})
}
