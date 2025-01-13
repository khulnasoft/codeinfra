package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func notImplemented(message string) codeinfra.AnyOutput {
	panic(message)
}

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		ctx.Export("result", codeinfra.Any(notImplemented("expression here is not implemented yet")))
		return nil
	})
}
