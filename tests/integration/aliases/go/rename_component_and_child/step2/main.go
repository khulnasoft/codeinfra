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

// Scenario #5 - composing #1 and #3 and making both changes at the same time
func NewFooComponent(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent43", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	parentOpt := codeinfra.Parent(fooComp)
	alias := &codeinfra.Alias{
		Name:   codeinfra.StringInput(codeinfra.String("otherchild")),
		Parent: fooComp,
	}
	aliasOpt := codeinfra.Aliases([]codeinfra.Alias{*alias})
	_, err = NewFooResource(ctx, "otherchildrenamed", parentOpt, aliasOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		alias := &codeinfra.Alias{Name: codeinfra.StringInput(codeinfra.String("comp5"))}
		aliasOpt := codeinfra.Aliases([]codeinfra.Alias{*alias})
		_, err := NewFooComponent(ctx, "newcomp5", aliasOpt)
		if err != nil {
			return err
		}

		return nil
	})
}
