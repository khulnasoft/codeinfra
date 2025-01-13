package main

import (
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		parent, err := NewRandom(ctx, "parent", &RandomArgs{
			Length: codeinfra.Int(8),
		})
		if err != nil {
			return err
		}

		child, err := NewRandom(ctx, "child", &RandomArgs{
			Length: codeinfra.Int(4),
		}, codeinfra.Parent(parent))
		if err != nil {
			return err
		}

		ctx.Export("parent", parent.Result)
		ctx.Export("child", child.Result)
		return nil
	})
}
