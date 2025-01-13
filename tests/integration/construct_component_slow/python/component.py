# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

from typing import Any, Optional

import codeinfra

class Component(codeinfra.ComponentResource):
    def __init__(self, name: str, opts: Optional[codeinfra.ResourceOptions] = None):
        props = dict()
        super().__init__("testcomponent:index:Component", name, props, opts, True)
