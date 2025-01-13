import codeinfra
import codeinfra_simple_invoke as simple_invoke

explicit_provider = simple_invoke.Provider("explicitProvider")
first = simple_invoke.StringResource("first", text="first hello")
data = simple_invoke.my_invoke_output(value="hello", opts=codeinfra.InvokeOutputOptions(depends_on=[first]))
second = simple_invoke.StringResource("second", text=data.result)
codeinfra.export("hello", data.result)
