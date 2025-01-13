package main

import (
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws"
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/ec2"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		// Create a new security group for port 80.
		securityGroup, err := ec2.NewSecurityGroup(ctx, "securityGroup", &ec2.SecurityGroupArgs{
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					Protocol: codeinfra.String("tcp"),
					FromPort: codeinfra.Int(0),
					ToPort:   codeinfra.Int(0),
					CidrBlocks: codeinfra.StringArray{
						codeinfra.String("0.0.0.0/0"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		// Get the ID for the latest Amazon Linux AMI.
		ami, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
			Filters: []aws.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"amzn-ami-hvm-*-x86_64-ebs",
					},
				},
			},
			Owners: []string{
				"137112412989",
			},
			MostRecent: codeinfra.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}
		// Create a simple web server using the startup script for the instance.
		server, err := ec2.NewInstance(ctx, "server", &ec2.InstanceArgs{
			Tags: codeinfra.StringMap{
				"Name": codeinfra.String("web-server-www"),
			},
			InstanceType: codeinfra.String(ec2.InstanceType_T2_Micro),
			SecurityGroups: codeinfra.StringArray{
				securityGroup.Name,
			},
			Ami:      codeinfra.String(ami.Id),
			UserData: codeinfra.String("#!/bin/bash\necho \"Hello, World!\" > index.html\nnohup python -m SimpleHTTPServer 80 &\n"),
		})
		if err != nil {
			return err
		}
		ctx.Export("publicIp", server.PublicIp)
		ctx.Export("publicHostName", server.PublicDns)
		return nil
	})
}
