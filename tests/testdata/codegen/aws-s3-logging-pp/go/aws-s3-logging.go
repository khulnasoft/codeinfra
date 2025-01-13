package main

import (
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/s3"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		logs, err := s3.NewBucket(ctx, "logs", nil)
		if err != nil {
			return err
		}
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Loggings: s3.BucketLoggingArray{
				&s3.BucketLoggingArgs{
					TargetBucket: logs.Bucket,
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("targetBucket", bucket.Loggings.ApplyT(func(loggings []s3.BucketLogging) (*string, error) {
			return &loggings[0].TargetBucket, nil
		}).(codeinfra.StringPtrOutput))
		return nil
	})
}
