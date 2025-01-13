import codeinfra
import codeinfra_simple as simple

prov = simple.Provider("prov")
res = simple.Resource("res", value=True,
opts = codeinfra.ResourceOptions(provider=prov))
