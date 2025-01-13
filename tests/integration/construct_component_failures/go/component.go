// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type componentArgs struct {
	Foo string `codeinfra:"foo"`
}

type ComponentArgs struct {
	Foo codeinfra.StringInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type Component struct {
	codeinfra.ResourceState

	Foo codeinfra.StringOutput `codeinfra:"foo"`
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
