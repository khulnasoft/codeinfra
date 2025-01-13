import * as codeinfra from "@codeinfra/codeinfra";
import * as aws from "@codeinfra/aws";

const mybucket = new aws.s3.Bucket("mybucket", {website: {
    indexDocument: "index.html",
}});
const indexhtml = new aws.s3.BucketObject("indexhtml", {
    bucket: mybucket.id,
    source: new codeinfra.asset.StringAsset("<h1>Hello, world!</h1>"),
    acl: "public-read",
    contentType: "text/html",
});
export const bucketEndpoint = codeinfra.interpolate`http://${mybucket.websiteEndpoint}`;
