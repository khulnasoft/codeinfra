package main

import (
	"github.com/khulnasoft/codeinfra-gcp/sdk/v8/go/gcp/storage"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		// Create a GCP resource (Storage Bucket)
		bucket, err := storage.NewBucket(ctx, "my-bucket", &storage.BucketArgs{
			Location: codeinfra.String("US"),
		})
		if err != nil {
			return err
		}

		// Export the DNS name of the bucket
		ctx.Export("bucketName", bucket.Url)
		return nil
	})
}
