// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		if _, err := NewComponent(ctx, "component", &ComponentArgs{
			Foo: &FooArgs{
				Something: codeinfra.String("hello"),
			},
			Bar: &BarArgs{
				Tags: codeinfra.StringMap{
					"a": codeinfra.String("world"),
					"b": codeinfra.ToSecret("shh").(codeinfra.StringOutput),
				},
			},
		}); err != nil {
			return err
		}
		return nil
	})
}
