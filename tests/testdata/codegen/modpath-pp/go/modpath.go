package main

import (
	other "git.example.org/thirdparty/sdk/go/pkg"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := other.NewThing(ctx, "thing", &other.ThingArgs{
			Idea: codeinfra.String("myIdea"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
