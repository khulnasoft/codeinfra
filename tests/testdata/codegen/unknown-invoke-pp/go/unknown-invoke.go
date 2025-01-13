package main

import (
	"github.com/khulnasoft/codeinfra-unknown/sdk/go/unknown"
	"github.com/khulnasoft/codeinfra-unknown/sdk/go/unknown/eks"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		data, err := unknown.GetData(ctx, map[string]interface{}{
			"input": "hello",
		}, nil)
		if err != nil {
			return err
		}
		_, err = eks.ModuleValues(ctx, map[string]interface{}{}, nil)
		if err != nil {
			return err
		}
		ctx.Export("content", data.Content)
		return nil
	})
}
