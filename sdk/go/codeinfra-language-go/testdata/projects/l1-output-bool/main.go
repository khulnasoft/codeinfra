package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		ctx.Export("output_true", codeinfra.Bool(true))
		ctx.Export("output_false", codeinfra.Bool(false))
		return nil
	})
}
