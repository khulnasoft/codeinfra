// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		component, err := NewComponent(ctx, "component")
		if err != nil {
			return err
		}
		result, err := component.CreateRandom(ctx, &ComponentCreateRandomArgs{
			Length: codeinfra.Int(10),
		})
		if err != nil {
			return err
		}
		ctx.Export("result", result.Result())
		return nil
	})
}
