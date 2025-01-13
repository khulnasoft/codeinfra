import asyncio
import codeinfra

output = codeinfra.Output.from_input(asyncio.sleep(3, []))
output.apply(lambda x: x[0])

foo = codeinfra.Output.from_input(asyncio.sleep(1, "foo"))
foo.apply(print)

printed = codeinfra.Output.from_input(asyncio.sleep(2, "printed"))
printed.apply(print)

not_printed = codeinfra.Output.from_input(asyncio.sleep(4, "not printed"))
not_printed.apply(print)
