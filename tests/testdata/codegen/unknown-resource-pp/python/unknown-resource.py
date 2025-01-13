import codeinfra
import codeinfra_unknown as unknown

provider = codeinfra.providers.Unknown("provider")
main = unknown.index.Main("main",
    first=hello,
    second={
        foo: bar,
    })
from_module = []
for range in [{"value": i} for i in range(0, 10)]:
    from_module.append(unknown.eks.Example(f"fromModule-{range['value']}", associated_main=main.id))
codeinfra.export("mainId", main["id"])
codeinfra.export("values", from_module["values"]["first"])
