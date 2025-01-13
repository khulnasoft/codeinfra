import codeinfra

ref = codeinfra.StackReference("ref", stack_name="organization/other/dev")
codeinfra.export("plain", ref.get_output("plain"))
codeinfra.export("secret", ref.get_output("secret"))
