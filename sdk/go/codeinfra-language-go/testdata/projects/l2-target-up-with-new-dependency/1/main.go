package main

import (
	"example.com/codeinfra-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := simple.NewResource(ctx, "targetOnly", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		})
		if err != nil {
			return err
		}
		dep, err := simple.NewResource(ctx, "dep", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "unrelated", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		}, codeinfra.DependsOn([]codeinfra.Resource{
			dep,
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
