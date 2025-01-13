// Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

// FooComponent is a component resource
type FooComponent struct {
	codeinfra.ResourceState
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		fooComponent := &FooComponent{}
		alias := &codeinfra.Alias{
			Name: codeinfra.String("foo"),
		}
		opts := codeinfra.Aliases([]codeinfra.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})
}
