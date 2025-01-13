// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type Component struct {
	codeinfra.ResourceState
}

func NewComponent(ctx *codeinfra.Context, name string, args ComponentArgs,
	opts ...codeinfra.ResourceOption,
) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, &args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type componentArgs struct {
	Children *int `codeinfra:"children"`
}

type ComponentArgs struct {
	Children *int
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		children := 5
		if _, err := NewComponent(ctx, "component", ComponentArgs{Children: &children}); err != nil {
			return err
		}
		return nil
	})
}
