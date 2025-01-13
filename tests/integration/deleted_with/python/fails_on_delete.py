# Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

from typing import Optional

import codeinfra

class FailsOnDelete(codeinfra.CustomResource):
    def __init__(self,
                 resource_name: str,
                 opts: Optional[codeinfra.ResourceOptions] = None):
        super().__init__("testprovider:index:FailsOnDelete", resource_name, {}, opts)
