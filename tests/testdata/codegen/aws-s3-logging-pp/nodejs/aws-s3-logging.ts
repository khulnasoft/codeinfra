import * as codeinfra from "@codeinfra/codeinfra";
import * as aws from "@codeinfra/aws";

const logs = new aws.s3.Bucket("logs", {});
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});
export const targetBucket = bucket.loggings.apply(loggings => loggings?.[0]?.targetBucket);
