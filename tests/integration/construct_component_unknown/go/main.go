// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"fmt"

	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		r, err := NewRandom(ctx, "resource", &RandomArgs{Length: codeinfra.Int(10)})
		if err != nil {
			return err
		}
		_, err = NewComponent(ctx, "component", &ComponentArgs{
			Message: r.ID().ApplyT(func(id codeinfra.ID) string {
				return fmt.Sprintf("message %v", id)
			}).(codeinfra.StringOutput),
			Nested: &ComponentNestedArgs{
				Value: r.ID().ApplyT(func(id codeinfra.ID) string {
					return fmt.Sprintf("nested.value %v", id)
				}).(codeinfra.StringOutput),
			},
		})
		return err
	})
}
