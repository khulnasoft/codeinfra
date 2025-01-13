"""A Google Cloud Python Codeinfra program"""

import codeinfra
from codeinfra_gcp import storage

# Create a GCP resource (Storage Bucket)
bucket = storage.Bucket('my-bucket', location="US")

# Export the DNS name of the bucket
codeinfra.export('bucket_name', bucket.url)
