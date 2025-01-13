package main

import (
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra/config"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		cfg := config.New(ctx, "")
		configLexicalName := cfg.Require("cC-Charlie_charlie.😃⁉️")
		resourceLexicalName, err := random.NewRandomPet(ctx, "aA-Alpha_alpha.🤯⁉️", &random.RandomPetArgs{
			Prefix: codeinfra.String(configLexicalName),
		})
		if err != nil {
			return err
		}
		ctx.Export("bB-Beta_beta.💜⁉", resourceLexicalName.ID())
		ctx.Export("dD-Delta_delta.🔥⁉", resourceLexicalName.ID())
		return nil
	})
}
