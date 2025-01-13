# Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import codeinfra


class FooResource(codeinfra.ComponentResource):
    def __init__(self, name, opts=None):
        alias_opts = codeinfra.ResourceOptions(aliases=[codeinfra.Alias(type_="my:module:FooResource")])
        opts = codeinfra.ResourceOptions.merge(opts, alias_opts)
        super().__init__("my:module:FooResourceNew", name, None, opts)


class ComponentResource(codeinfra.ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentResource", name, None, opts)
        FooResource("child", codeinfra.ResourceOptions(parent=self))


ComponentResource("comp")
