# Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

from typing import Optional

import codeinfra

class Component(codeinfra.ComponentResource):
    foo: codeinfra.Output[str]

    def __init__(self, name: str, foo: codeinfra.Input[str], opts: Optional[codeinfra.ResourceOptions] = None):
        props = dict()
        props["foo"] = foo
        super().__init__("testcomponent:index:Component", name, props, opts, True)

    def get_message(self) -> codeinfra.Output[str]:
        __args__ = dict()
        __args__['__self__'] = self
        return codeinfra.runtime.call('testcomponent:index:Component/getMessage', __args__, res=self)
