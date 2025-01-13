package main

import (
	"github.com/khulnasoft/codeinfra-aws/sdk/v4/go/aws"
	"github.com/khulnasoft/codeinfra-aws/sdk/v4/go/aws/s3"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		provider, err := aws.NewProvider(ctx, "provider", &aws.ProviderArgs{
			Region: codeinfra.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucket(ctx, "bucket1", nil, codeinfra.Provider(provider), codeinfra.DependsOn([]codeinfra.Resource{
			provider,
		}), codeinfra.Protect(true), codeinfra.IgnoreChanges([]string{
			"bucket",
			"lifecycleRules[0]",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
