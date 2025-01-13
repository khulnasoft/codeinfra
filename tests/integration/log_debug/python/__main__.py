# Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

import codeinfra

class MyComponent(codeinfra.ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("test:index:MyComponent", name, {}, opts)

codeinfra.log.debug("A debug message")

c = MyComponent("mycomponent")
