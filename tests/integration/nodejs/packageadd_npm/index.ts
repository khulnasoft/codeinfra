import * as codeinfra from "@codeinfra/codeinfra";
import * as gcp from "@codeinfra/gcp";

// Create a GCP resource (Storage Bucket)
const bucket = new gcp.storage.Bucket("my-bucket", {
    location: "US"
});

// Export the DNS name of the bucket
export const bucketName = bucket.url;
