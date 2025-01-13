package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		ctx.Export("zero", codeinfra.Float64(0))
		ctx.Export("one", codeinfra.Float64(1))
		ctx.Export("e", codeinfra.Float64(2.718))
		ctx.Export("minInt32", codeinfra.Float64(-2147483648))
		ctx.Export("max", codeinfra.Float64(1.7976931348623157e+308))
		ctx.Export("min", codeinfra.Float64(5e-324))
		return nil
	})
}
