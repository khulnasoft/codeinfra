package main

import (
	"example.com/codeinfra-large/sdk/go/v4/large"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		res, err := large.NewString(ctx, "res", &large.StringArgs{
			Value: codeinfra.String("hello world"),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", res.Value)
		return nil
	})
}
