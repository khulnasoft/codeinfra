package main

import (
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/lambda"
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/s3"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		siteBucket, err := s3.NewBucket(ctx, "siteBucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "testFileAsset", &s3.BucketObjectArgs{
			Bucket: siteBucket.ID(),
			Source: codeinfra.NewFileAsset("file.txt"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "testStringAsset", &s3.BucketObjectArgs{
			Bucket: siteBucket.ID(),
			Source: codeinfra.NewStringAsset("<h1>File contents</h1>"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "testRemoteAsset", &s3.BucketObjectArgs{
			Bucket: siteBucket.ID(),
			Source: codeinfra.NewRemoteAsset("https://codeinfra.test"),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewFunction(ctx, "testFileArchive", &lambda.FunctionArgs{
			Role: siteBucket.Arn,
			Code: codeinfra.NewFileArchive("file.tar.gz"),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewFunction(ctx, "testRemoteArchive", &lambda.FunctionArgs{
			Role: siteBucket.Arn,
			Code: codeinfra.NewRemoteArchive("https://codeinfra.test/foo.tar.gz"),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewFunction(ctx, "testAssetArchive", &lambda.FunctionArgs{
			Role: siteBucket.Arn,
			Code: codeinfra.NewAssetArchive(map[string]interface{}{
				"file.txt":   codeinfra.NewFileAsset("file.txt"),
				"string.txt": codeinfra.NewStringAsset("<h1>File contents</h1>"),
				"remote.txt": codeinfra.NewRemoteAsset("https://codeinfra.test"),
				"file.tar":   codeinfra.NewFileArchive("file.tar.gz"),
				"remote.tar": codeinfra.NewRemoteArchive("https://codeinfra.test/foo.tar.gz"),
				".nestedDir": codeinfra.NewAssetArchive(map[string]interface{}{
					"file.txt":   codeinfra.NewFileAsset("file.txt"),
					"string.txt": codeinfra.NewStringAsset("<h1>File contents</h1>"),
					"remote.txt": codeinfra.NewRemoteAsset("https://codeinfra.test"),
					"file.tar":   codeinfra.NewFileArchive("file.tar.gz"),
					"remote.tar": codeinfra.NewRemoteArchive("https://codeinfra.test/foo.tar.gz"),
				}),
			}),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
