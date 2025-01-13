package main

import (
	"example.com/codeinfra-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		provider, err := simple.NewProvider(ctx, "provider", nil)
		if err != nil {
			return err
		}
		parent1, err := simple.NewResource(ctx, "parent1", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		}, codeinfra.Provider(provider))
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "child1", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		}, codeinfra.Parent(parent1))
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "orphan1", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		})
		if err != nil {
			return err
		}
		parent2, err := simple.NewResource(ctx, "parent2", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		}, codeinfra.Protect(true))
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "child2", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		}, codeinfra.Parent(parent2))
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "orphan2", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
