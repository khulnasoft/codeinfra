package main

import (
	"encoding/json"

	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/iam"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Effect":   "Allow",
					"Action":   "lambda:*",
					"Resource": "arn:aws:lambda:*:*:function:*",
					"Condition": map[string]interface{}{
						"StringEquals": map[string]interface{}{
							"aws:RequestTag/Team": []string{
								"iamuser-admin",
								"iamuser2-admin",
							},
						},
						"ForAllValues:StringEquals": map[string]interface{}{
							"aws:TagKeys": []string{
								"Team",
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		// Create a policy with multiple Condition keys
		policy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
			Path:        codeinfra.String("/"),
			Description: codeinfra.String("My test policy"),
			Policy:      codeinfra.String(json0),
		})
		if err != nil {
			return err
		}
		ctx.Export("policyName", policy.Name)
		return nil
	})
}
