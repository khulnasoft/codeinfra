# Copyright 2016-2022, KhulnaSoft Ltd.  All rights reserved.

from typing import Optional

import codeinfra


class Child(codeinfra.ComponentResource):
    @codeinfra.input_type
    class ChildArgs:
        pass

    def __init__(
        self,
        resource_name: str,
        message: Optional[str] = None,
        opts: Optional[codeinfra.ResourceOptions] = None,
    ):
        props = Container.ContainerArgs.__new__(Container.ContainerArgs)
        props.__dict__["message"] = message
        super().__init__("test:index:Child", resource_name, props, opts)
        if opts and opts.urn:
            return
        self.register_outputs({ "message": message })

    @property
    @codeinfra.getter
    def message(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "message")


class Container(codeinfra.ComponentResource):
    @codeinfra.input_type
    class ContainerArgs:
        pass

    def __init__(
        self,
        resource_name: str,
        child: Optional[Child] = None,
        opts: Optional[codeinfra.ResourceOptions] = None,
    ):
        props = Container.ContainerArgs.__new__(Container.ContainerArgs)
        props.__dict__["child"] = child
        super().__init__("test:index:Container", resource_name, props, opts)
        if opts and opts.urn:
            return
        self.register_outputs({ "child": child })

    @property
    @codeinfra.getter
    def child(self) -> codeinfra.Output[Child]:
        return codeinfra.get(self, "child")


class Module(codeinfra.runtime.ResourceModule):
    def version(self):
        return "0.0.1"

    def construct(self, name: str, typ: str, urn: str) -> codeinfra.Resource:
        if typ == "test:index:Child":
            return Child(name, opts=codeinfra.ResourceOptions(urn=urn))
        else:
            raise Exception(f"unknown resource type {typ}")


codeinfra.runtime.register_resource_module("test", "index", Module())


child = Child("mychild", message="hello world!")
container = Container("mycontainer", child=child)


def assert_equal(args):
    expected_urn = args["expected_urn"]
    actual_urn = args["actual_urn"]
    actual_message = args["actual_message"]
    assert expected_urn == actual_urn, \
        f"expected urn '{expected_urn}' not equal to actual urn '{actual_urn}'"
    assert "hello world!" == actual_message, \
        f"expected message 'hello world!' not equal to actual message '{actual_message}'"


def round_trip(urn: str):
    round_tripped_container = Container("mycontainer", opts=codeinfra.ResourceOptions(urn=urn))
    codeinfra.Output.all(
        expected_urn=child.urn,
        actual_urn=round_tripped_container.child.urn,
        actual_message=round_tripped_container.child.message,
    ).apply(assert_equal)


container.urn.apply(round_trip)
