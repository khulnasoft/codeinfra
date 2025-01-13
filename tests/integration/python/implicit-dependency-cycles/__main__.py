# Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

# Regression test for https://github.com/khulnasoft/codeinfra/issues/13551
import codeinfra

class A(codeinfra.ComponentResource):
    def __init__(self, name: str, opts=None):
        super().__init__("my:modules:A", name, {}, opts)
        a_1 = codeinfra.CustomResource("my:module:Child-1", "a-child-1", opts=codeinfra.ResourceOptions(parent=self, depends_on=[self]))
        codeinfra.CustomResource("my:module:Child-2", "a-child-2", props={"transitive_urn": a_1.urn} ,opts=codeinfra.ResourceOptions(parent=self))

A("a")
