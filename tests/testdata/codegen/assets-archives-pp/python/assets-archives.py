import codeinfra
import codeinfra_aws as aws

site_bucket = aws.s3.Bucket("siteBucket")
test_file_asset = aws.s3.BucketObject("testFileAsset",
    bucket=site_bucket.id,
    source=codeinfra.FileAsset("file.txt"))
test_string_asset = aws.s3.BucketObject("testStringAsset",
    bucket=site_bucket.id,
    source=codeinfra.StringAsset("<h1>File contents</h1>"))
test_remote_asset = aws.s3.BucketObject("testRemoteAsset",
    bucket=site_bucket.id,
    source=codeinfra.RemoteAsset("https://codeinfra.test"))
test_file_archive = aws.lambda_.Function("testFileArchive",
    role=site_bucket.arn,
    code=codeinfra.FileArchive("file.tar.gz"))
test_remote_archive = aws.lambda_.Function("testRemoteArchive",
    role=site_bucket.arn,
    code=codeinfra.RemoteArchive("https://codeinfra.test/foo.tar.gz"))
test_asset_archive = aws.lambda_.Function("testAssetArchive",
    role=site_bucket.arn,
    code=codeinfra.AssetArchive({
        "file.txt": codeinfra.FileAsset("file.txt"),
        "string.txt": codeinfra.StringAsset("<h1>File contents</h1>"),
        "remote.txt": codeinfra.RemoteAsset("https://codeinfra.test"),
        "file.tar": codeinfra.FileArchive("file.tar.gz"),
        "remote.tar": codeinfra.RemoteArchive("https://codeinfra.test/foo.tar.gz"),
        ".nestedDir": codeinfra.AssetArchive({
            "file.txt": codeinfra.FileAsset("file.txt"),
            "string.txt": codeinfra.StringAsset("<h1>File contents</h1>"),
            "remote.txt": codeinfra.RemoteAsset("https://codeinfra.test"),
            "file.tar": codeinfra.FileArchive("file.tar.gz"),
            "remote.tar": codeinfra.RemoteArchive("https://codeinfra.test/foo.tar.gz"),
        }),
    }))
