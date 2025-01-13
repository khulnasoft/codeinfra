//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		var x []string
		ctx.Export("a", codeinfra.String(x[0]))
		return nil
	})
}
