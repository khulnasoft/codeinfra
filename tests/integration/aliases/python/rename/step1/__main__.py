# Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

from codeinfra import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #1 - rename a resource
res1 = Resource1("res1")
