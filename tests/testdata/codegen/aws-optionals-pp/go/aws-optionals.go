package main

import (
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/iam"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		policyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid: codeinfra.StringRef("1"),
					Actions: []string{
						"s3:ListAllMyBuckets",
						"s3:GetBucketLocation",
					},
					Resources: []string{
						"arn:aws:s3:::*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewPolicy(ctx, "example", &iam.PolicyArgs{
			Name:   codeinfra.String("example_policy"),
			Path:   codeinfra.String("/"),
			Policy: codeinfra.String(policyDocument.Json),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
