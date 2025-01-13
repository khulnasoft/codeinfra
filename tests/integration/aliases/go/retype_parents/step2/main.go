// Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

type Resource struct {
	codeinfra.ResourceState
}

type ComponentSix struct {
	codeinfra.ResourceState
}

type ComponentSixParent struct {
	codeinfra.ResourceState
}

func NewResource(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*Resource, error) {
	comp := &Resource{}
	err := ctx.RegisterComponentResource("my:module:Resource", name, comp, opts...)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

// Scenario #6 - Nested parents changing types
func NewComponentSix(ctx *codeinfra.Context, name string, opts ...codeinfra.ResourceOption) (*ComponentSix, error) {
	comp := &ComponentSix{}
	aliases := make([]codeinfra.Alias, 0)
	for i := 0; i < 100; i++ {
		aliases = append(aliases, codeinfra.Alias{
			Type: codeinfra.Sprintf("my:module:ComponentSix-v%d", i),
		})
	}
	err := ctx.RegisterComponentResource(
		"my:module:ComponentSix-v0", name, comp,
		codeinfra.Aliases(aliases), codeinfra.Composite(opts...))
	if err != nil {
		return nil, err
	}
	parentOpt := codeinfra.Parent(comp)
	_, err = NewResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

func NewComponentSixParent(ctx *codeinfra.Context, name string,
	opts ...codeinfra.ResourceOption,
) (*ComponentSixParent, error) {
	comp := &ComponentSixParent{}
	aliases := make([]codeinfra.Alias, 0)
	for i := 0; i < 10; i++ {
		aliases = append(aliases, codeinfra.Alias{
			Type: codeinfra.Sprintf("my:module:ComponentSixParent-v%d", i),
		})
	}
	err := ctx.RegisterComponentResource(
		"my:module:ComponentSixParent-v0", name, comp,
		codeinfra.Aliases(aliases), codeinfra.Composite(opts...))
	if err != nil {
		return nil, err
	}
	parentOpt := codeinfra.Parent(comp)
	_, err = NewComponentSix(ctx, "child", parentOpt)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := NewComponentSixParent(ctx, "comp6")
		if err != nil {
			return err
		}

		return nil
	})
}
