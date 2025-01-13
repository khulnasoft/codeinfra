import codeinfra
import codeinfra_simple_invoke as simple_invoke

codeinfra.export("hello", simple_invoke.my_invoke_output(value="hello").apply(lambda invoke: invoke.result))
codeinfra.export("goodbye", simple_invoke.my_invoke_output(value="goodbye").apply(lambda invoke: invoke.result))
