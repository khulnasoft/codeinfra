# Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.

"""An example program that should be Pylint clean"""

import binascii
import os
import codeinfra
from codeinfra.dynamic import Resource, ResourceProvider, CreateResult


class RandomResourceProvider(ResourceProvider):
    """Random resource provider."""

    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, {"val": val})


class Random(Resource):
    """Random resource."""
    val: str

    def __init__(self, name, opts=None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)


r = Random("foo")

codeinfra.export("cwd", os.getcwd())
codeinfra.export("random_urn", r.urn)
codeinfra.export("random_id", r.id)
codeinfra.export("random_val", r.val)
