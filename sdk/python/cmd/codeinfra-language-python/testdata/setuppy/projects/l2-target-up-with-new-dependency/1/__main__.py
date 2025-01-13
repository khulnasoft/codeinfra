import codeinfra
import codeinfra_simple as simple

target_only = simple.Resource("targetOnly", value=True)
dep = simple.Resource("dep", value=True)
unrelated = simple.Resource("unrelated", value=True,
opts = codeinfra.ResourceOptions(depends_on=[dep]))
