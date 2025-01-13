# Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.

from typing import Any, Optional

import codeinfra

class Component(codeinfra.ComponentResource):
    echo: codeinfra.Output[Any]
    childId: codeinfra.Output[str]

    def __init__(self, name: str, echo: codeinfra.Input[Any], opts: Optional[codeinfra.ResourceOptions] = None):
        props = dict()
        props["echo"] = echo
        props["childId"] = None
        props["secret"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)
