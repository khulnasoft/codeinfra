import asyncio
import codeinfra


async def raises():
    await asyncio.sleep(0)
    raise Exception("oh no")


async def catches():
    try:
        await raises()
    except Exception:
        return "oh yeah"

output = codeinfra.Output.from_input(catches())
output.apply(print)
