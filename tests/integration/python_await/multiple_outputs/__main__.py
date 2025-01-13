import asyncio
import codeinfra

output = codeinfra.Output.from_input(asyncio.sleep(3, "magic string"))
output.apply(print)

exported = codeinfra.Output.from_input(asyncio.sleep(2, "foo"))
codeinfra.export("exported", exported)
exported.apply(print)

another = codeinfra.Output.from_input(asyncio.sleep(5, "bar"))
another.apply(print)


