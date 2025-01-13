package main

import (
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		randomPassword, err := random.NewRandomPassword(ctx, "randomPassword", &random.RandomPasswordArgs{
			Length:          codeinfra.Int(16),
			Special:         codeinfra.Bool(true),
			OverrideSpecial: codeinfra.String("_%@"),
		})
		if err != nil {
			return err
		}
		ctx.Export("password", randomPassword.Result)
		return nil
	})
}
