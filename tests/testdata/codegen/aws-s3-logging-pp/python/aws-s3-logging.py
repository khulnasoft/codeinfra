import codeinfra
import codeinfra_aws as aws

logs = aws.s3.Bucket("logs")
bucket = aws.s3.Bucket("bucket", loggings=[{
    "target_bucket": logs.bucket,
}])
codeinfra.export("targetBucket", bucket.loggings[0].target_bucket)
