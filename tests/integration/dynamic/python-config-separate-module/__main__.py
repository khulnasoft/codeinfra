# Copyright 2024, KhulnaSoft Ltd.  All rights reserved.

import codeinfra
from codeinfra.dynamic import Resource

from provider import SimpleProvider


class SimpleResource(Resource):
    authenticated: codeinfra.Output[str]
    color: codeinfra.Output[str]

    def __init__(self, name):
        super().__init__(SimpleProvider(), name, {"authenticated": None, "color": None})


r = SimpleResource("foo")
codeinfra.export("authenticated", r.authenticated)
codeinfra.export("color", r.color)
