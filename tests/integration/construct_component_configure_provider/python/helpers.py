import asyncio
import codeinfra


def unknownIfDryRun(value):
    if codeinfra.runtime.is_dry_run():
        return codeinfra.Output(resources=set(), future=fut(None), is_known=fut(False))
    return codeinfra.Output.from_input(value)


def fut(x):
    f = asyncio.Future()
    f.set_result(x)
    return f
