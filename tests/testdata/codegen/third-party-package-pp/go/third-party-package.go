package main

import (
	other "git.example.org/thirdparty/sdk/go/pkg"
	"git.example.org/thirdparty/sdk/go/pkg/module"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := other.NewThing(ctx, "Other", &other.ThingArgs{
			Idea: codeinfra.String("Support Third Party"),
		})
		if err != nil {
			return err
		}
		_, err = module.NewObject(ctx, "Question", &module.ObjectArgs{
			Answer: codeinfra.Float64(42),
		})
		if err != nil {
			return err
		}
		_, err = module.NewObject(ctx, "Question2", &module.ObjectArgs{
			Answer: codeinfra.Float64(24),
		})
		if err != nil {
			return err
		}
		_, err = other.NewProvider(ctx, "Provider", &other.ProviderArgs{
			ObjectProp: codeinfra.StringMap{
				"prop1": codeinfra.String("foo"),
				"prop2": codeinfra.String("bar"),
				"prop3": codeinfra.String("fizz"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
