import codeinfra
import codeinfra_unknown as unknown

data = unknown.index.get_data(input="hello")
values = unknown.eks.module_values()
codeinfra.export("content", data["content"])
