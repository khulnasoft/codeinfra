# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

from typing import Optional

import codeinfra

class Component(codeinfra.ComponentResource):
    def __init__(self,
                 name: str,
                 opts: Optional[codeinfra.ResourceOptions] = None):
        super().__init__("testcomponent:index:Component", name, {}, opts, True)

    @codeinfra.output_type
    class GetMessageResult:
        def __init__(self, message: str):
            if message and not isinstance(message, str):
                raise TypeError("Expected argument 'message' to be a str")
            codeinfra.set(self, "message", message)

        @property
        @codeinfra.getter
        def message(self) -> str:
            return codeinfra.get(self, "message")

    def get_message(__self__, echo: codeinfra.Input[str]) -> codeinfra.Output['Component.GetMessageResult']:
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['echo'] = echo
        return codeinfra.runtime.call('testcomponent:index:Component/getMessage',
                                   __args__,
                                   res=__self__,
                                   typ=Component.GetMessageResult)
