# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

from typing import Optional

import codeinfra

class Random(codeinfra.CustomResource):
    def __init__(self,
                 resource_name: str,
                 length: codeinfra.Input[int],
                 opts: Optional[codeinfra.ResourceOptions] = None):
        props = {
            "length": length,
            "result": None,
        }
        super().__init__("testprovider:index:Random", resource_name, props, opts)

    @property
    @codeinfra.getter
    def length(self) -> codeinfra.Output[int]:
        return codeinfra.get(self, "length")

    @property
    @codeinfra.getter
    def result(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "result")
