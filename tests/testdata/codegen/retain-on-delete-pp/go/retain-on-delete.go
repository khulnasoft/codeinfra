package main

import (
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := random.NewRandomPet(ctx, "foo", nil, codeinfra.RetainOnDelete(true))
		if err != nil {
			return err
		}
		return nil
	})
}
