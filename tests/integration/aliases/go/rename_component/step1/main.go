// Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

// FooComponent is a component resource
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

// Scenario #3 - rename a component (and all it's children)
func NewFooComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent42", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	// Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
	// alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
	parentOpt := codeinfra.Parent(fooComp)
	_, err = NewFooResource(ctx, name+"-child", parentOpt)
	if err != nil {
		return nil, err
	}
	_, err = NewFooResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := NewFooComponent(ctx, "comp3")
		if err != nil {
			return err
		}

		return nil
	})
}
