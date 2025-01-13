# Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import codeinfra

from component import Component
from test_provider import TestProvider

test_provider = TestProvider("testProvider")

component1 = Component("component1", first="Hello", second="World",
    opts=codeinfra.ResourceOptions(provider=test_provider))
result1 = component1.get_message("Alice")

component2 = Component("component2", first="Hi", second="There",
    opts=codeinfra.ResourceOptions(providers=[test_provider]))
result2 = component2.get_message("Bob")


codeinfra.export("message1", result1.message)
codeinfra.export("message2", result2.message)
