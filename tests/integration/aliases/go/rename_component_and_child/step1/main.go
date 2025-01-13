// Copyright 2016-2022, KhulnaSoft Ltd.  All rights reserved.
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

// Scenario #5 - composing #1 and #3 and making both changes at the same time
func NewFooComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent43", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	parentOpt := codeinfra.Parent(fooComp)
	_, err = NewFooResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := NewFooComponent(ctx, "comp5")
		if err != nil {
			return err
		}

		return nil
	})
}
