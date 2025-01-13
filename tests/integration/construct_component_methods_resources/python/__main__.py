# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import codeinfra

from component import Component

component = Component("component")
result = component.create_random(length=10).result

codeinfra.export("result", result)
