package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		ctx.Export("empty", codeinfra.Map{})
		ctx.Export("strings", codeinfra.StringMap{
			"greeting": codeinfra.String("Hello, world!"),
			"farewell": codeinfra.String("Goodbye, world!"),
		})
		ctx.Export("numbers", codeinfra.Float64Map{
			"1": codeinfra.Float64(1),
			"2": codeinfra.Float64(2),
		})
		ctx.Export("keys", codeinfra.Float64Map{
			"my.key": codeinfra.Float64(1),
			"my-key": codeinfra.Float64(2),
			"my_key": codeinfra.Float64(3),
			"MY_KEY": codeinfra.Float64(4),
			"mykey":  codeinfra.Float64(5),
			"MYKEY":  codeinfra.Float64(6),
		})
		return nil
	})
}
