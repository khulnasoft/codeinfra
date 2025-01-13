# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

from typing import Optional

import codeinfra


class Provider(codeinfra.ProviderResource):
    message: codeinfra.Output[str]

    def __init__(self, name: str, message: codeinfra.Input[str], opts: Optional[codeinfra.ResourceOptions] = None) -> None:
        super().__init__("testcomponent", name, {"message": message}, opts)


class Component(codeinfra.ComponentResource):
    message: codeinfra.Output[str]

    def __init__(self, name: str, opts: Optional[codeinfra.ResourceOptions] = None) -> None:
        props = {
            "message": None
        }
        super().__init__("testcomponent:index:Component", name, props, opts, True)


component = Component("mycomponent", codeinfra.ResourceOptions(
    providers={
        "testcomponent": Provider("myprovider", "hello world"),
    })
)


codeinfra.export("message", component.message)
