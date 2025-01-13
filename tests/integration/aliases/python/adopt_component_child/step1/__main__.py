# Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import codeinfra


class FooResource(codeinfra.ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:FooResource", name, None, opts)


class ComponentResource(codeinfra.ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentResource", name, None, opts)


ComponentResource("comp")


FooResource("child")
