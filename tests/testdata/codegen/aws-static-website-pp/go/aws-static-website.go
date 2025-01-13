package main

import (
	awsstaticwebsite "github.com/khulnasoft/codeinfra-aws-static-website/sdk/go/aws-static-website"
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/cloudfront"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := awsstaticwebsite.NewWebsite(ctx, "websiteResource", &awsstaticwebsite.WebsiteArgs{
			SitePath:  codeinfra.String("string"),
			IndexHTML: codeinfra.String("string"),
			CacheTTL:  codeinfra.Float64(0),
			CdnArgs: &awsstaticwebsite.CDNArgsArgs{
				CloudfrontFunctionAssociations: cloudfront.DistributionOrderedCacheBehaviorFunctionAssociationArray{
					&cloudfront.DistributionOrderedCacheBehaviorFunctionAssociationArgs{
						EventType:   codeinfra.String("string"),
						FunctionArn: codeinfra.String("string"),
					},
				},
				ForwardedValues: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesArgs{
					Cookies: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs{
						Forward: codeinfra.String("string"),
						WhitelistedNames: codeinfra.StringArray{
							codeinfra.String("string"),
						},
					},
					QueryString: codeinfra.Bool(false),
					Headers: codeinfra.StringArray{
						codeinfra.String("string"),
					},
					QueryStringCacheKeys: codeinfra.StringArray{
						codeinfra.String("string"),
					},
				},
				LambdaFunctionAssociations: cloudfront.DistributionOrderedCacheBehaviorLambdaFunctionAssociationArray{
					&cloudfront.DistributionOrderedCacheBehaviorLambdaFunctionAssociationArgs{
						EventType:   codeinfra.String("string"),
						LambdaArn:   codeinfra.String("string"),
						IncludeBody: codeinfra.Bool(false),
					},
				},
			},
			CertificateARN:          codeinfra.String("string"),
			Error404:                codeinfra.String("string"),
			AddWebsiteVersionHeader: codeinfra.Bool(false),
			PriceClass:              codeinfra.String("string"),
			AtomicDeployments:       codeinfra.Bool(false),
			Subdomain:               codeinfra.String("string"),
			TargetDomain:            codeinfra.String("string"),
			WithCDN:                 codeinfra.Bool(false),
			WithLogs:                codeinfra.Bool(false),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
