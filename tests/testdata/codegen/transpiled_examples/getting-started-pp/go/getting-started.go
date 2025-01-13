package main

import (
	"fmt"

	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/s3"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		mybucket, err := s3.NewBucket(ctx, "mybucket", &s3.BucketArgs{
			Website: &s3.BucketWebsiteArgs{
				IndexDocument: codeinfra.String("index.html"),
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "indexhtml", &s3.BucketObjectArgs{
			Bucket:      mybucket.ID(),
			Source:      codeinfra.NewStringAsset("<h1>Hello, world!</h1>"),
			Acl:         codeinfra.String("public-read"),
			ContentType: codeinfra.String("text/html"),
		})
		if err != nil {
			return err
		}
		ctx.Export("bucketEndpoint", mybucket.WebsiteEndpoint.ApplyT(func(websiteEndpoint string) (string, error) {
			return fmt.Sprintf("http://%v", websiteEndpoint), nil
		}).(codeinfra.StringOutput))
		return nil
	})
}
