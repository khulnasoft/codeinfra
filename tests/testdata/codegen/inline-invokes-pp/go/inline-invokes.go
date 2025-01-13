package main

import (
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/ec2"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		invokeLookupVpc, err := ec2.LookupVpc(ctx, &ec2.LookupVpcArgs{
			Default: codeinfra.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSecurityGroup(ctx, "webSecurityGroup", &ec2.SecurityGroupArgs{
			VpcId: invokeLookupVpc.Id,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
