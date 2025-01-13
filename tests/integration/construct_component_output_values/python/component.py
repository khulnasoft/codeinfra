# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import codeinfra
from typing import Mapping, Optional


@codeinfra.input_type
class BarArgs:
    def __init__(__self__, *,
                 tags: Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]] = None):
        if tags is not None:
            codeinfra.set(__self__, "tags", tags)

    @property
    @codeinfra.getter
    def tags(self) -> Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]]:
        return codeinfra.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]]):
        codeinfra.set(self, "tags", value)


@codeinfra.input_type
class FooArgs:
    def __init__(__self__, *,
                 something: Optional[codeinfra.Input[str]] = None):
        if something is not None:
            codeinfra.set(__self__, "something", something)

    @property
    @codeinfra.getter
    def something(self) -> Optional[codeinfra.Input[str]]:
        return codeinfra.get(self, "something")

    @something.setter
    def something(self, value: Optional[codeinfra.Input[str]]):
        codeinfra.set(self, "something", value)


@codeinfra.input_type
class ComponentArgs:
    def __init__(__self__, *,
                 bar: Optional[codeinfra.Input['BarArgs']] = None,
                 foo: Optional['FooArgs'] = None):
        if bar is not None:
            codeinfra.set(__self__, "bar", bar)
        if foo is not None:
            codeinfra.set(__self__, "foo", foo)

    @property
    @codeinfra.getter
    def bar(self) -> Optional[codeinfra.Input['BarArgs']]:
        return codeinfra.get(self, "bar")

    @bar.setter
    def bar(self, value: Optional[codeinfra.Input['BarArgs']]):
        codeinfra.set(self, "bar", value)

    @property
    @codeinfra.getter
    def foo(self) -> Optional['FooArgs']:
        return codeinfra.get(self, "foo")

    @foo.setter
    def foo(self, value: Optional['FooArgs']):
        codeinfra.set(self, "foo", value)


class Component(codeinfra.ComponentResource):
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ComponentArgs] = None,
                 opts: Optional[codeinfra.ResourceOptions] = None):
        super().__init__('testcomponent:index:Component', resource_name, args, opts, remote=True)
