package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		ref, err := codeinfra.NewStackReference(ctx, "ref", &codeinfra.StackReferenceArgs{
			Name: codeinfra.String("organization/other/dev"),
		})
		if err != nil {
			return err
		}
		ctx.Export("plain", ref.GetOutput(codeinfra.String("plain")))
		ctx.Export("secret", ref.GetOutput(codeinfra.String("secret")))
		return nil
	})
}
