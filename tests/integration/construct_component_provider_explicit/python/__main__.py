# Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

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


class LocalComponent(codeinfra.ComponentResource):
    message: codeinfra.Output[str]

    def __init__(self, name: str, opts: Optional[codeinfra.ResourceOptions] = None) -> None:
        super().__init__("my:index:LocalComponent", name, {}, opts)

        component = Component(f"{name}-mycomponent", codeinfra.ResourceOptions(parent=self))
        self.message = component.message


provider = Provider("myprovider", "hello world")
component = Component("mycomponent", codeinfra.ResourceOptions(
    provider=provider,
))
localComponent = LocalComponent("mylocalcomponent", codeinfra.ResourceOptions(
    providers=[provider],
))

codeinfra.export("message", component.message)
codeinfra.export("nestedMessage", localComponent.message)
