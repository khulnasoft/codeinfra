package main

import (
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/ec2"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		aws_vpc, err := ec2.NewVpc(ctx, "aws_vpc", &ec2.VpcArgs{
			CidrBlock:       codeinfra.String("10.0.0.0/16"),
			InstanceTenancy: codeinfra.String("default"),
		})
		if err != nil {
			return err
		}
		privateS3VpcEndpoint, err := ec2.NewVpcEndpoint(ctx, "privateS3VpcEndpoint", &ec2.VpcEndpointArgs{
			VpcId:       aws_vpc.ID(),
			ServiceName: codeinfra.String("com.amazonaws.us-west-2.s3"),
		})
		if err != nil {
			return err
		}
		privateS3PrefixList := ec2.GetPrefixListOutput(ctx, ec2.GetPrefixListOutputArgs{
			PrefixListId: privateS3VpcEndpoint.PrefixListId,
		}, nil)
		bar, err := ec2.NewNetworkAcl(ctx, "bar", &ec2.NetworkAclArgs{
			VpcId: aws_vpc.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewNetworkAclRule(ctx, "privateS3NetworkAclRule", &ec2.NetworkAclRuleArgs{
			NetworkAclId: bar.ID(),
			RuleNumber:   codeinfra.Int(200),
			Egress:       codeinfra.Bool(false),
			Protocol:     codeinfra.String("tcp"),
			RuleAction:   codeinfra.String("allow"),
			CidrBlock: privateS3PrefixList.ApplyT(func(privateS3PrefixList ec2.GetPrefixListResult) (string, error) {
				return privateS3PrefixList.CidrBlocks[0], nil
			}).(codeinfra.StringOutput),
			FromPort: codeinfra.Int(443),
			ToPort:   codeinfra.Int(443),
		})
		if err != nil {
			return err
		}
		// A contrived example to test that helper nested records ( `filters`
		// below) generate correctly when using output-versioned function
		// invoke forms.
		_ = ec2.GetAmiIdsOutput(ctx, ec2.GetAmiIdsOutputArgs{
			Owners: codeinfra.StringArray{
				bar.ID(),
			},
			Filters: ec2.GetAmiIdsFilterArray{
				&ec2.GetAmiIdsFilterArgs{
					Name: bar.ID(),
					Values: codeinfra.StringArray{
						codeinfra.String("codeinfra*"),
					},
				},
			},
		}, nil)
		return nil
	})
}
