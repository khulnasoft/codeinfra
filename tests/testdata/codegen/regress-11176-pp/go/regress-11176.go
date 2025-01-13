package main

import (
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/ecs"
	awsxecs "github.com/khulnasoft/codeinfra-awsx/sdk/go/awsx/ecs"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		cluster, err := ecs.NewCluster(ctx, "cluster", nil)
		if err != nil {
			return err
		}
		_, err = awsxecs.NewFargateService(ctx, "nginx", &awsxecs.FargateServiceArgs{
			Cluster: cluster.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
