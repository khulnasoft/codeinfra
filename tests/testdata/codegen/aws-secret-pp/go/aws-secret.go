package main

import (
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/rds"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{
			MasterPassword: codeinfra.ToSecret("foobar").(codeinfra.StringOutput),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
