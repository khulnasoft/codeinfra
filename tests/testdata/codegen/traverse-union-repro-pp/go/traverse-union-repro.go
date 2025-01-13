package main

import (
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/fsx"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := fsx.NewOpenZfsFileSystem(ctx, "test", &fsx.OpenZfsFileSystemArgs{
			StorageCapacity: codeinfra.Int(64),
			SubnetIds: codeinfra.String{
				aws_subnet.Test1.Id,
			},
			DeploymentType:     codeinfra.String("SINGLE_AZ_1"),
			ThroughputCapacity: codeinfra.Int(64),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
