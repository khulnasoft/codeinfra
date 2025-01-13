package main

import (
	"fmt"

	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/s3"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		myBucket, err := s3.NewBucket(ctx, "myBucket", &s3.BucketArgs{
			Website: &s3.BucketWebsiteArgs{
				IndexDocument: codeinfra.String("index.html"),
			},
		})
		if err != nil {
			return err
		}
		ownershipControls, err := s3.NewBucketOwnershipControls(ctx, "ownershipControls", &s3.BucketOwnershipControlsArgs{
			Bucket: myBucket.ID(),
			Rule: &s3.BucketOwnershipControlsRuleArgs{
				ObjectOwnership: codeinfra.String("ObjectWriter"),
			},
		})
		if err != nil {
			return err
		}
		publicAccessBlock, err := s3.NewBucketPublicAccessBlock(ctx, "publicAccessBlock", &s3.BucketPublicAccessBlockArgs{
			Bucket:          myBucket.ID(),
			BlockPublicAcls: codeinfra.Bool(false),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "index.html", &s3.BucketObjectArgs{
			Bucket:      myBucket.ID(),
			Source:      codeinfra.NewFileAsset("./index.html"),
			ContentType: codeinfra.String("text/html"),
			Acl:         codeinfra.String("public-read"),
		}, codeinfra.DependsOn([]codeinfra.Resource{
			publicAccessBlock,
			ownershipControls,
		}))
		if err != nil {
			return err
		}
		ctx.Export("bucketName", myBucket.ID())
		ctx.Export("bucketEndpoint", myBucket.WebsiteEndpoint.ApplyT(func(websiteEndpoint string) (string, error) {
			return fmt.Sprintf("http://%v", websiteEndpoint), nil
		}).(codeinfra.StringOutput))
		return nil
	})
}
