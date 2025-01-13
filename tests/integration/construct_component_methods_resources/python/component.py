# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

from typing import Optional

import codeinfra

class Component(codeinfra.ComponentResource):
    def __init__(self,
                 name: str,
                 opts: Optional[codeinfra.ResourceOptions] = None):
        super().__init__("testcomponent:index:Component", name, {}, opts, True)

    @codeinfra.output_type
    class CreateRandomResult:
        def __init__(self, result: str):
            if result and not isinstance(result, str):
                raise TypeError("Expected argument 'result' to be a str")
            codeinfra.set(self, "result", result)

        @property
        @codeinfra.getter
        def result(self) -> str:
            return codeinfra.get(self, "result")

    def create_random(__self__, length: codeinfra.Input[int]) -> codeinfra.Output['Component.CreateRandomResult']:
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['length'] = length
        return codeinfra.runtime.call('testcomponent:index:Component/createRandom',
                                   __args__,
                                   res=__self__,
                                   typ=Component.CreateRandomResult)
