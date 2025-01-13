import codeinfra
import codeinfra_simple as simple
import codeinfra_simple_invoke as simple_invoke

res = simple.Resource("res", value=True)
codeinfra.export("nonSecret", simple_invoke.secret_invoke_output(value="hello",
    secret_response=False).apply(lambda invoke: invoke.response))
codeinfra.export("firstSecret", simple_invoke.secret_invoke_output(value="hello",
    secret_response=res.value).apply(lambda invoke: invoke.response))
codeinfra.export("secondSecret", simple_invoke.secret_invoke_output(value=codeinfra.Output.secret("goodbye"),
    secret_response=False).apply(lambda invoke: invoke.response))
