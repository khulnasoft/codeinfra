# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

from codeinfra import Output

from component import Component, ComponentArgs, FooArgs, BarArgs

Component("component", ComponentArgs(
    foo=FooArgs(something="hello"),
    bar=BarArgs(tags={
        "a": "world",
        "b": Output.secret("shh"),
    })
))
