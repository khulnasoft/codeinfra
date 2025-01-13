# Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

from typing import Any, Optional

import codeinfra

class Component(codeinfra.ComponentResource):
    def __init__(self, name: str, id: codeinfra.Input[str], opts: Optional[codeinfra.ResourceOptions] = None):
        props = dict()
        props["id"] = id
        super().__init__("testcomponent:index:Component", name, props, opts, True)

    @property
    @codeinfra.getter
    def id(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "id")