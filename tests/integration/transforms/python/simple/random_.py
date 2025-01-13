# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

from typing import Optional

import codeinfra

class Random(codeinfra.CustomResource):
    def __init__(self,
                 resource_name: str,
                 length: codeinfra.Input[int],
                 prefix: Optional[codeinfra.Input[str]] = None,
                 opts: Optional[codeinfra.ResourceOptions] = None):
        props = {
            "length": length,
            "result": None,
            "prefix": prefix,
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

    def invoke(self, args):
        return codeinfra.runtime.invoke("testprovider:index:returnArgs", args)


class Component(codeinfra.ComponentResource):
    def __init__(self,
                 resource_name: str,
                 length: codeinfra.Input[int],
                 opts: Optional[codeinfra.ResourceOptions] = None):
        props = {
            "length": length,
            "childId": None,
        }
        super().__init__("testprovider:index:Component", resource_name, props, opts, True)

    @property
    @codeinfra.getter
    def length(self) -> codeinfra.Output[int]:
        return codeinfra.get(self, "length")

    @property
    @codeinfra.getter
    def child_id(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "childId")

class Provider(codeinfra.ProviderResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[codeinfra.ResourceOptions] = None):
        super(Provider, __self__).__init__(
            'testprovider',
            resource_name,
            None,
            opts)
