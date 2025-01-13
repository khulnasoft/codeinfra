import codeinfra
import codeinfra_goodbye as goodbye

prov = goodbye.Provider("prov", text="World")
# The resource name is based on the parameter value
res = goodbye.Goodbye("res", opts = codeinfra.ResourceOptions(provider=prov))
codeinfra.export("parameterValue", res.parameter_value)
