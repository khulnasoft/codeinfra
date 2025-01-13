//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		ctx.Export("aboutName", codeinfra.String("about-name"))
		return nil
	})
}
