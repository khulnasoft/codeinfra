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

type FooComponent2 struct {
	codeinfra.ResourceState
}

type FooComponent3 struct {
	codeinfra.ResourceState
}

type FooComponent4 struct {
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
	var nilInput codeinfra.StringInput
	aliasURN := codeinfra.CreateURN(
		codeinfra.StringInput(codeinfra.String("res2")),
		codeinfra.StringInput(codeinfra.String("my:module:FooResource")),
		nilInput,
		codeinfra.StringInput(codeinfra.String(ctx.Project())),
		codeinfra.StringInput(codeinfra.String(ctx.Stack())))
	alias := &codeinfra.Alias{
		URN: aliasURN,
	}
	aliasOpt := codeinfra.Aliases([]codeinfra.Alias{*alias})
	parentOpt := codeinfra.Parent(fooComp)
	_, err = NewFooResource(ctx, name+"-child", aliasOpt, parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent2(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*FooComponent2, error) {
	fooComp := &FooComponent2{}
	err := ctx.RegisterComponentResource("my:module:FooComponent2", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent3(ctx *codeinfra.Context,
	name string,
	childAliasParent codeinfra.Resource,
	opts ...codeinfra.ResourceOption,
) (*FooComponent3, error) {
	fooComp := &FooComponent3{}
	err := ctx.RegisterComponentResource("my:module:FooComponent3", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}

	alias := &codeinfra.Alias{}
	if childAliasParent != nil {
		alias.Parent = childAliasParent
	} else {
		alias.NoParent = codeinfra.Bool(true)
	}

	aliasOpt := codeinfra.Aliases([]codeinfra.Alias{*alias})
	parentOpt := codeinfra.Parent(fooComp)
	_, err = NewFooComponent2(ctx, name+"-child", aliasOpt, parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent4(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*FooComponent4, error) {
	fooComp := &FooComponent4{}
	alias := &codeinfra.Alias{
		Parent: nil,
	}
	aliasOpt := codeinfra.Aliases([]codeinfra.Alias{*alias, *alias})
	o := []codeinfra.ResourceOption{aliasOpt}
	o = append(o, opts...)
	err := ctx.RegisterComponentResource("my:module:FooComponent4", name, fooComp, o...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		comp2, err := NewFooComponent(ctx, "comp2")
		if err != nil {
			return err
		}
		alias := &codeinfra.Alias{
			NoParent: codeinfra.Bool(true),
		}
		aliasOpt := codeinfra.Aliases([]codeinfra.Alias{*alias})
		parentOpt := codeinfra.Parent(comp2)
		_, err = NewFooComponent2(ctx, "unparented", aliasOpt, parentOpt)
		if err != nil {
			return err
		}
		_, err = NewFooComponent3(ctx, "parentedbystack", nil)
		if err != nil {
			return err
		}
		pbcOpt := codeinfra.Parent(comp2)
		_, err = NewFooComponent3(ctx, "parentedbycomponent", comp2, pbcOpt)
		if err != nil {
			return err
		}
		dupeOpt := codeinfra.Parent(comp2)
		_, err = NewFooComponent4(ctx, "duplicateAliases", dupeOpt)
		if err != nil {
			return err
		}
		return nil
	})
}
