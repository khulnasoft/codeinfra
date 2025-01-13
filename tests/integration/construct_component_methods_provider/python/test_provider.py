# Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import codeinfra

class TestProvider(codeinfra.ProviderResource):
    def __init__(__self__, resource_name: str):
        super().__init__("testprovider", resource_name)
