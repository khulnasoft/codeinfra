import codeinfra

stack_ref = codeinfra.StackReference("stackRef", stack_name="PLACEHOLDER_ORG_NAME/stackreference-producer/PLACEHOLDER_STACK_NAME")
codeinfra.export("referencedImageName", stack_ref.outputs["imageName"])
