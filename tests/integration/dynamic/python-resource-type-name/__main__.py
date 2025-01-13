# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

from codeinfra import export
from codeinfra.dynamic import CreateResult, Resource, ResourceProvider


class CustomResource(
    Resource, module="custom-provider", name="CustomResource"
):
    def __init__(self, name, opts=None):
        super().__init__(DummyResourceProvider(), name, {}, opts)


class DummyResourceProvider(ResourceProvider):
    def create(self, props):
        return CreateResult("resource-id", {})


resource = CustomResource("resource-name")

export("urn", resource.urn)
