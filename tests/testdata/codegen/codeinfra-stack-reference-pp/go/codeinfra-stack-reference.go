package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := codeinfra.NewStackReference(ctx, "stackRef", &codeinfra.StackReferenceArgs{
			Name: codeinfra.String("foo/bar/dev"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
