# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

from typing import Optional

import codeinfra

class Component(codeinfra.ComponentResource):
    def __init__(self, name: str, args: codeinfra.Inputs, opts: Optional[codeinfra.ResourceOptions] = None):
        super().__init__("testcomponent:index:Component", name, args, opts, True)
