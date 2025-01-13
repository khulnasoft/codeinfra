import codeinfra
import codeinfra_large as large

res = large.String("res", value="hello world")
codeinfra.export("output", res.value)
