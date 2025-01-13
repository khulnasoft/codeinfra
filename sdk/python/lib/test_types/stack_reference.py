import codeinfra

ref = codeinfra.StackReference("ref")
ref.outputs["bucket"].apply(lambda bucket: codeinfra.log.info(f"Bucket: {bucket}"))
