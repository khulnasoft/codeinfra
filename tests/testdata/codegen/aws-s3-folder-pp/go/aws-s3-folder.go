package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/iam"
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/s3"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		// Create a bucket and expose a website index document
		siteBucket, err := s3.NewBucket(ctx, "siteBucket", &s3.BucketArgs{
			Website: &s3.BucketWebsiteArgs{
				IndexDocument: codeinfra.String("index.html"),
			},
		})
		if err != nil {
			return err
		}
		siteDir := "www"
		// For each file in the directory, create an S3 object stored in `siteBucket`
		files0, err := os.ReadDir(siteDir)
		if err != nil {
			return err
		}
		fileNames0 := make([]string, len(files0))
		for key0, val0 := range files0 {
			fileNames0[key0] = val0.Name()
		}
		var files []*s3.BucketObject
		for key0, val0 := range fileNames0 {
			__res, err := s3.NewBucketObject(ctx, fmt.Sprintf("files-%v", key0), &s3.BucketObjectArgs{
				Bucket:      siteBucket.ID(),
				Key:         codeinfra.String(val0),
				Source:      codeinfra.NewFileAsset(fmt.Sprintf("%v/%v", siteDir, val0)),
				ContentType: codeinfra.String(val0),
			}, codeinfra.DeletedWith(siteBucket))
			if err != nil {
				return err
			}
			files = append(files, __res)
		}
		// Set the access policy for the bucket so all objects are readable
		_, err = s3.NewBucketPolicy(ctx, "bucketPolicy", &s3.BucketPolicyArgs{
			Bucket: siteBucket.ID(),
			Policy: siteBucket.ID().ApplyT(func(id string) (codeinfra.String, error) {
				var _zero codeinfra.String
				tmpJSON0, err := json.Marshal(map[string]interface{}{
					"Version": "2012-10-17",
					"Statement": []map[string]interface{}{
						map[string]interface{}{
							"Effect":    "Allow",
							"Principal": "*",
							"Action": []string{
								"s3:GetObject",
							},
							"Resource": []string{
								fmt.Sprintf("arn:aws:s3:::%v/*", id),
							},
						},
					},
				})
				if err != nil {
					return _zero, err
				}
				json0 := string(tmpJSON0)
				return codeinfra.String(json0), nil
			}).(codeinfra.StringOutput),
		})
		if err != nil {
			return err
		}
		ctx.Export("bucketName", siteBucket.Bucket)
		ctx.Export("websiteUrl", siteBucket.WebsiteEndpoint)
		return nil
	})
}
