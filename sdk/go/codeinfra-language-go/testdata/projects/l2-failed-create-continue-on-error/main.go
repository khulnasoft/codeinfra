package main

import (
	"example.com/codeinfra-fail_on_create/sdk/go/v4/fail_on_create"
	"example.com/codeinfra-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		failing, err := fail_on_create.NewResource(ctx, "failing", &fail_on_create.ResourceArgs{
			Value: codeinfra.Bool(false),
		})
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "dependent", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		}, codeinfra.DependsOn([]codeinfra.Resource{
			failing,
		}))
		if err != nil {
			return err
		}
		dependent_on_output, err := simple.NewResource(ctx, "dependent_on_output", &simple.ResourceArgs{
			Value: failing.Value,
		})
		if err != nil {
			return err
		}
		independent, err := simple.NewResource(ctx, "independent", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "double_dependency", &simple.ResourceArgs{
			Value: codeinfra.Bool(true),
		}, codeinfra.DependsOn([]codeinfra.Resource{
			independent,
			dependent_on_output,
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
