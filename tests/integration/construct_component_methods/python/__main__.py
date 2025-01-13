# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import codeinfra

from component import Component

component = Component("component", first="Hello", second="World")
result = component.get_message("Alice")

codeinfra.export("message", result.message)
