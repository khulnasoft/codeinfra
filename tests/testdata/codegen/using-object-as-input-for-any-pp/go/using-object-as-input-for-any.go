package main

import (
	"github.com/khulnasoft/codeinfra-aws-native/sdk/go/aws/iam"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			RoleName: codeinfra.String("ScriptIAMRole"),
			AssumeRolePolicyDocument: codeinfra.Any(map[string]interface{}{
				"Version": "2012-10-17",
				"Statement": []map[string]interface{}{
					map[string]interface{}{
						"Effect": "Allow",
						"Action": "sts:AssumeRole",
						"Principal": map[string]interface{}{
							"Service": []string{
								"cloudformation.amazonaws.com",
								"gamelift.amazonaws.com",
							},
						},
					},
				},
			}),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
