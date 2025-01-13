import asyncio
import binascii
import os
from codeinfra import Output
from codeinfra.dynamic import Resource, ResourceProvider, CreateResult


class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, {"val": val})


class Random(Resource):
    val: str

    def __init__(self, name, opts=None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)


r = Random("foo")

output = Output.from_input(asyncio.sleep(2, "magic_string"))
