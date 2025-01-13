import codeinfra
import codeinfra_subpackage as subpackage

# The resource name is based on the parameter value
example = subpackage.HelloWorld("example")
codeinfra.export("parameterValue", example.parameter_value)
