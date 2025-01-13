// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type FooResource struct {
	codeinfra.ResourceState
}

type FooComponent struct {
	codeinfra.ResourceState
}

func NewFooResource(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}
	return fooRes, nil
}

func NewFooComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		foo, err := NewFooComponent(ctx, "foo")
		if err != nil {
			return err
		}

		aliasOpt := codeinfra.Aliases([]codeinfra.Alias{{
			Parent: foo,
		}})
		_, err = NewFooResource(ctx, "child", aliasOpt)
		if err != nil {
			return err
		}

		return nil
	})
}
