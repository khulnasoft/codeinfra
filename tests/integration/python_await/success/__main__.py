import asyncio
import codeinfra

output = codeinfra.Output.from_input(asyncio.sleep(1, "magic string"))
output.apply(print)
