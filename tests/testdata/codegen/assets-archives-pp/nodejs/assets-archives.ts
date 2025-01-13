import * as codeinfra from "@codeinfra/codeinfra";
import * as aws from "@codeinfra/aws";

const siteBucket = new aws.s3.Bucket("siteBucket", {});
const testFileAsset = new aws.s3.BucketObject("testFileAsset", {
    bucket: siteBucket.id,
    source: new codeinfra.asset.FileAsset("file.txt"),
});
const testStringAsset = new aws.s3.BucketObject("testStringAsset", {
    bucket: siteBucket.id,
    source: new codeinfra.asset.StringAsset("<h1>File contents</h1>"),
});
const testRemoteAsset = new aws.s3.BucketObject("testRemoteAsset", {
    bucket: siteBucket.id,
    source: new codeinfra.asset.RemoteAsset("https://codeinfra.test"),
});
const testFileArchive = new aws.lambda.Function("testFileArchive", {
    role: siteBucket.arn,
    code: new codeinfra.asset.FileArchive("file.tar.gz"),
});
const testRemoteArchive = new aws.lambda.Function("testRemoteArchive", {
    role: siteBucket.arn,
    code: new codeinfra.asset.RemoteArchive("https://codeinfra.test/foo.tar.gz"),
});
const testAssetArchive = new aws.lambda.Function("testAssetArchive", {
    role: siteBucket.arn,
    code: new codeinfra.asset.AssetArchive({
        "file.txt": new codeinfra.asset.FileAsset("file.txt"),
        "string.txt": new codeinfra.asset.StringAsset("<h1>File contents</h1>"),
        "remote.txt": new codeinfra.asset.RemoteAsset("https://codeinfra.test"),
        "file.tar": new codeinfra.asset.FileArchive("file.tar.gz"),
        "remote.tar": new codeinfra.asset.RemoteArchive("https://codeinfra.test/foo.tar.gz"),
        ".nestedDir": new codeinfra.asset.AssetArchive({
            "file.txt": new codeinfra.asset.FileAsset("file.txt"),
            "string.txt": new codeinfra.asset.StringAsset("<h1>File contents</h1>"),
            "remote.txt": new codeinfra.asset.RemoteAsset("https://codeinfra.test"),
            "file.tar": new codeinfra.asset.FileArchive("file.tar.gz"),
            "remote.tar": new codeinfra.asset.RemoteArchive("https://codeinfra.test/foo.tar.gz"),
        }),
    }),
});
