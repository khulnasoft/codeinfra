# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import sys

import codeinfra

from component import Component
from random_ import Random


def panic(text: str):
    print(text)
    sys.exit(1)


r = Random("resource", length=10)
component = Component("component")

result = component.get_message(r.id)

codeinfra.export("result", result.apply(lambda v: panic("should not run (result)")))
