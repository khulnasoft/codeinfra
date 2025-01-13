import asyncio
import codeinfra

output = codeinfra.Output.from_input(asyncio.sleep(1, []))
output.apply(lambda x: x[0])
