// Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.
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

// Scenario #4 - change the type of a component
func NewFooComponent(ctx *codeinfra.Context, name string) (*FooComponent, error) {
	fooComp := &FooComponent{}
	alias := &codeinfra.Alias{
		Type: codeinfra.StringInput(codeinfra.String("my:module:FooComponent44")),
	}
	aliasOpt := codeinfra.Aliases([]codeinfra.Alias{*alias})
	err := ctx.RegisterComponentResource("my:diffmodule:FooComponent55DiffType", name, fooComp, aliasOpt)
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
		_, err := NewFooComponent(ctx, "comp4")
		if err != nil {
			return err
		}

		return nil
	})
}
