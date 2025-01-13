# Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import codeinfra

from typing import Mapping, Optional, Sequence


CAMEL_TO_SNAKE_CASE_TABLE = {
    "anotherValue": "another_value",
    "listOfNestedValues": "list_of_nested_values",
    "nestedValue": "nested_value",
    "serviceName": "service_name",
    "somethingElse": "something_else",
    "someValue": "some_value",
}

SNAKE_TO_CAMEL_CASE_TABLE = {
    "another_value": "anotherValue",
    "list_of_nested_values": "listOfNestedValues",
    "nested_value": "nestedValue",
    "service_name": "serviceName",
    "something_else": "somethingElse",
    "some_value": "someValue",
}


@codeinfra.input_type
class NestedArgs:
    def __init__(self,
                 foo_bar: codeinfra.Input[str],
                 some_value: codeinfra.Input[int]):
        codeinfra.set(self, "foo_bar", foo_bar)
        codeinfra.set(self, "some_value", some_value)

    @property
    @codeinfra.getter(name="fooBar")
    def foo_bar(self) -> codeinfra.Input[str]:
        return codeinfra.get(self, "foo_bar")

    @property
    @codeinfra.getter(name="someValue")
    def some_value(self) -> codeinfra.Input[int]:
        return codeinfra.get(self, "some_value")


@codeinfra.output_type
class NestedOldOutput(dict):
    def __init__(self,
                 foo_bar: str,
                 some_value: int):
        codeinfra.set(self, "foo_bar", foo_bar)
        codeinfra.set(self, "some_value", some_value)

    @property
    @codeinfra.getter(name="fooBar")
    def foo_bar(self) -> str:
        return codeinfra.get(self, "foo_bar")

    @property
    @codeinfra.getter(name="someValue")
    def some_value(self) -> int:
        return codeinfra.get(self, "some_value")

    def _translate_property(self, prop):
        return CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


class OldBehavior(codeinfra.CustomResource):
    def __init__(self,
                 resource_name: str,
                 service_name: codeinfra.Input[str],
                 some_value: Optional[codeinfra.Input[str]] = None,
                 another_value: Optional[codeinfra.Input[str]] = None,
                 nested_value: Optional[codeinfra.Input[codeinfra.InputType[NestedArgs]]] = None,
                 list_of_nested_values: Optional[codeinfra.Input[Sequence[codeinfra.Input[codeinfra.InputType[NestedArgs]]]]] = None,
                 tags: Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]] = None,
                 items: Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]] = None,
                 keys: Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]] = None,
                 values: Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]] = None,
                 get: Optional[codeinfra.Input[str]] = None,
                 opts: Optional[codeinfra.ResourceOptions] = None):
        if opts is None:
            opts = codeinfra.ResourceOptions()
        __props__ = dict()
        __props__["service_name"] = service_name
        __props__["some_value"] = some_value
        __props__["another_value"] = another_value
        __props__["nested_value"] = nested_value
        __props__["list_of_nested_values"] = list_of_nested_values
        __props__["tags"] = tags
        __props__["items"] = items
        __props__["keys"] = keys
        __props__["values"] = values
        __props__["get"] = get
        __props__["something_else"] = None
        super().__init__("test::OldBehavior", resource_name, __props__, opts)

    @property
    @codeinfra.getter(name="serviceName")
    def service_name(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "service_name")

    @property
    @codeinfra.getter(name="someValue")
    def some_value(self) -> codeinfra.Output[Optional[str]]:
        return codeinfra.get(self, "some_value")

    @property
    @codeinfra.getter(name="anotherValue")
    def another_value(self) -> codeinfra.Output[Optional[str]]:
        return codeinfra.get(self, "another_value")

    @property
    @codeinfra.getter(name="nestedValue")
    def nested_value(self) -> codeinfra.Output[Optional[NestedOldOutput]]:
        return codeinfra.get(self, "nested_value")

    @property
    @codeinfra.getter(name="listOfNestedValues")
    def list_of_nested_values(self) -> codeinfra.Output[Optional[Sequence[NestedOldOutput]]]:
        return codeinfra.get(self, "list_of_nested_values")

    @property
    @codeinfra.getter
    def tags(self) -> codeinfra.Output[Optional[Mapping[str, str]]]:
        return codeinfra.get(self, "tags")

    @property
    @codeinfra.getter
    def items(self) -> codeinfra.Output[Optional[Mapping[str, str]]]:
        return codeinfra.get(self, "items")

    @property
    @codeinfra.getter
    def keys(self) -> codeinfra.Output[Optional[Sequence[str]]]:
        return codeinfra.get(self, "keys")

    @property
    @codeinfra.getter
    def values(self) -> codeinfra.Output[Optional[Sequence[str]]]:
        return codeinfra.get(self, "values")

    @property
    @codeinfra.getter
    def get(self) -> codeinfra.Output[Optional[str]]:
        return codeinfra.get(self, "get")

    @property
    @codeinfra.getter(name="somethingElse")
    def something_else(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "something_else")

    def translate_output_property(self, prop):
        return CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# NestedNewOutput doesn't have a _translate_property method.
@codeinfra.output_type
class NestedNewOutput(dict):
    def __init__(self,
                 foo_bar: str,
                 some_value: int):
        codeinfra.set(self, "foo_bar", foo_bar)
        codeinfra.set(self, "some_value", some_value)

    @property
    @codeinfra.getter(name="fooBar")
    def foo_bar(self) -> str:
        return codeinfra.get(self, "foo_bar")

    @property
    @codeinfra.getter(name="someValue")
    def some_value(self) -> int:
        return codeinfra.get(self, "some_value")


# The <Resource>Args class that is passed-through. Since this is an @input_type, its type/name
# metadata will be used for translations.
@codeinfra.input_type
class NewBehaviorArgs:
    def __init__(self,
                 service_name: codeinfra.Input[str],
                 some_value: Optional[codeinfra.Input[str]] = None,
                 another_value: Optional[codeinfra.Input[str]] = None,
                 nested_value: Optional[codeinfra.Input[codeinfra.InputType[NestedArgs]]] = None,
                 list_of_nested_values: Optional[codeinfra.Input[Sequence[codeinfra.Input[codeinfra.InputType[NestedArgs]]]]] = None,
                 tags: Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]] = None,
                 items: Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]] = None,
                 keys: Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]] = None,
                 values: Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]] = None,
                 get: Optional[codeinfra.Input[str]] = None):
        codeinfra.set(self, "service_name", service_name)
        codeinfra.set(self, "some_value", some_value)
        codeinfra.set(self, "another_value", another_value)
        codeinfra.set(self, "nested_value", nested_value)
        codeinfra.set(self, "list_of_nested_values", list_of_nested_values)
        codeinfra.set(self, "tags", tags)
        codeinfra.set(self, "items", items)
        codeinfra.set(self, "keys", keys)
        codeinfra.set(self, "values", values)
        codeinfra.set(self, "get", get)

    @property
    @codeinfra.getter(name="serviceName")
    def service_name(self) -> codeinfra.Input[str]:
        return codeinfra.get(self, "service_name")

    @property
    @codeinfra.getter(name="someValue")
    def some_value(self) -> Optional[codeinfra.Input[str]]:
        return codeinfra.get(self, "some_value")

    @property
    @codeinfra.getter(name="anotherValue")
    def another_value(self) -> Optional[codeinfra.Input[str]]:
        return codeinfra.get(self, "another_value")

    @property
    @codeinfra.getter(name="nestedValue")
    def nested_value(self) -> Optional[codeinfra.Input[codeinfra.InputType[NestedArgs]]]:
        return codeinfra.get(self, "nested_value")

    @property
    @codeinfra.getter(name="listOfNestedValues")
    def list_of_nested_values(self) -> Optional[codeinfra.Input[Sequence[codeinfra.Input[codeinfra.InputType[NestedArgs]]]]]:
        return codeinfra.get(self, "list_of_nested_values")

    @property
    @codeinfra.getter
    def tags(self) -> Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]]:
        return codeinfra.get(self, "tags")

    @property
    @codeinfra.getter
    def items(self) -> Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]]:
        return codeinfra.get(self, "items")

    @property
    @codeinfra.getter
    def keys(self) -> Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]]:
        return codeinfra.get(self, "keys")

    @property
    @codeinfra.getter
    def values(self) -> Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]]:
        return codeinfra.get(self, "values")

    @property
    @codeinfra.getter
    def get(self) -> Optional[codeinfra.Input[str]]:
        return codeinfra.get(self, "get")

# This resource opts-in to the new translation behavior, by passing its <Resource>Args type.
# It's written in the same manner that our codegen is emitted.
# The class includes overrides of `translate_output_property` and `translate_input_property`
# that raise exceptions, to ensure they aren't called.
class NewBehavior(codeinfra.CustomResource):
    def __init__(self,
                 resource_name: str,
                 service_name: codeinfra.Input[str],
                 some_value: Optional[codeinfra.Input[str]] = None,
                 another_value: Optional[codeinfra.Input[str]] = None,
                 nested_value: Optional[codeinfra.Input[codeinfra.InputType[NestedArgs]]] = None,
                 list_of_nested_values: Optional[codeinfra.Input[Sequence[codeinfra.Input[codeinfra.InputType[NestedArgs]]]]] = None,
                 tags: Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]] = None,
                 items: Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]] = None,
                 keys: Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]] = None,
                 values: Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]] = None,
                 get: Optional[codeinfra.Input[str]] = None,
                 opts: Optional[codeinfra.ResourceOptions] = None):
        if opts is None:
            opts = codeinfra.ResourceOptions()
        # This is how our codegen creates the props to pass to super's `__init__`.
        __props__ = NewBehaviorArgs.__new__(NewBehaviorArgs)
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["some_value"] = some_value
        __props__.__dict__["another_value"] = another_value
        __props__.__dict__["nested_value"] = nested_value
        __props__.__dict__["list_of_nested_values"] = list_of_nested_values
        __props__.__dict__["tags"] = tags
        __props__.__dict__["items"] = items
        __props__.__dict__["keys"] = keys
        __props__.__dict__["values"] = values
        __props__.__dict__["get"] = get
        __props__.__dict__["something_else"] = None
        super().__init__("test::NewBehavior", resource_name, __props__, opts)

    @property
    @codeinfra.getter(name="serviceName")
    def service_name(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "service_name")

    @property
    @codeinfra.getter(name="someValue")
    def some_value(self) -> codeinfra.Output[Optional[str]]:
        return codeinfra.get(self, "some_value")

    @property
    @codeinfra.getter(name="anotherValue")
    def another_value(self) -> codeinfra.Output[Optional[str]]:
        return codeinfra.get(self, "another_value")

    @property
    @codeinfra.getter(name="nestedValue")
    def nested_value(self) -> codeinfra.Output[Optional[NestedNewOutput]]:
        return codeinfra.get(self, "nested_value")

    @property
    @codeinfra.getter(name="listOfNestedValues")
    def list_of_nested_values(self) -> codeinfra.Output[Optional[Sequence[NestedNewOutput]]]:
        return codeinfra.get(self, "list_of_nested_values")

    @property
    @codeinfra.getter
    def tags(self) -> codeinfra.Output[Optional[Mapping[str, str]]]:
        return codeinfra.get(self, "tags")

    @property
    @codeinfra.getter
    def items(self) -> codeinfra.Output[Optional[Mapping[str, str]]]:
        return codeinfra.get(self, "items")

    @property
    @codeinfra.getter
    def keys(self) -> codeinfra.Output[Optional[Sequence[str]]]:
        return codeinfra.get(self, "keys")

    @property
    @codeinfra.getter
    def values(self) -> codeinfra.Output[Optional[Sequence[str]]]:
        return codeinfra.get(self, "values")

    @property
    @codeinfra.getter
    def get(self) -> codeinfra.Output[Optional[str]]:
        return codeinfra.get(self, "get")

    @property
    @codeinfra.getter(name="somethingElse")
    def something_else(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "something_else")

    def translate_output_property(self, prop):
        raise Exception("Unexpected call to translate_output_property")

    def translate_input_property(self, prop):
        raise Exception("Unexpected call to translate_input_property")


# Variant where the <Resource>Args class is a subclass of dict.
@codeinfra.input_type
class NewBehaviorDictArgs(dict):
    def __init__(self,
                 service_name: codeinfra.Input[str],
                 some_value: Optional[codeinfra.Input[str]] = None,
                 another_value: Optional[codeinfra.Input[str]] = None,
                 nested_value: Optional[codeinfra.Input[codeinfra.InputType[NestedArgs]]] = None,
                 list_of_nested_values: Optional[codeinfra.Input[Sequence[codeinfra.Input[codeinfra.InputType[NestedArgs]]]]] = None,
                 tags: Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]] = None,
                 items: Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]] = None,
                 keys: Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]] = None,
                 values: Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]] = None,
                 get: Optional[codeinfra.Input[str]] = None):
        codeinfra.set(self, "service_name", service_name)
        codeinfra.set(self, "some_value", some_value)
        codeinfra.set(self, "another_value", another_value)
        codeinfra.set(self, "nested_value", nested_value)
        codeinfra.set(self, "list_of_nested_values", list_of_nested_values)
        codeinfra.set(self, "tags", tags)
        codeinfra.set(self, "items", items)
        codeinfra.set(self, "keys", keys)
        codeinfra.set(self, "values", values)
        codeinfra.set(self, "get", get)

    @property
    @codeinfra.getter(name="serviceName")
    def service_name(self) -> codeinfra.Input[str]:
        return codeinfra.get(self, "service_name")

    @property
    @codeinfra.getter(name="someValue")
    def some_value(self) -> Optional[codeinfra.Input[str]]:
        return codeinfra.get(self, "some_value")

    @property
    @codeinfra.getter(name="anotherValue")
    def another_value(self) -> Optional[codeinfra.Input[str]]:
        return codeinfra.get(self, "another_value")

    @property
    @codeinfra.getter(name="nestedValue")
    def nested_value(self) -> Optional[codeinfra.Input[codeinfra.InputType[NestedArgs]]]:
        return codeinfra.get(self, "nested_value")

    @property
    @codeinfra.getter(name="listOfNestedValues")
    def list_of_nested_values(self) -> Optional[codeinfra.Input[Sequence[codeinfra.Input[codeinfra.InputType[NestedArgs]]]]]:
        return codeinfra.get(self, "list_of_nested_values")

    @property
    @codeinfra.getter
    def tags(self) -> Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]]:
        return codeinfra.get(self, "tags")

    @property
    @codeinfra.getter
    def items(self) -> Optional[codeinfra.Input[Mapping[str, codeinfra.Input[str]]]]:
        return codeinfra.get(self, "items")

    @property
    @codeinfra.getter
    def keys(self) -> Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]]:
        return codeinfra.get(self, "keys")

    @property
    @codeinfra.getter
    def values(self) -> Optional[codeinfra.Input[Sequence[codeinfra.Input[str]]]]:
        return codeinfra.get(self, "values")

    @property
    @codeinfra.getter
    def get(self) -> Optional[codeinfra.Input[str]]:
        return codeinfra.get(self, "get")


# Passes the <Resource>Args class (that's a subclass of dict) through.
class NewBehaviorDict(codeinfra.CustomResource):
    def __init__(self,
                 resource_name: str,
                 args: NewBehaviorDictArgs,
                 opts: Optional[codeinfra.ResourceOptions] = None):
        if opts is None:
            opts = codeinfra.ResourceOptions()

        # Include the additional output and pass args directly to super's `__init__`.
        args["something_else"] = None
        super().__init__("test::NewBehaviorDict", resource_name, args, opts)

    @property
    @codeinfra.getter(name="serviceName")
    def service_name(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "service_name")

    @property
    @codeinfra.getter(name="someValue")
    def some_value(self) -> codeinfra.Output[Optional[str]]:
        return codeinfra.get(self, "some_value")

    @property
    @codeinfra.getter(name="anotherValue")
    def another_value(self) -> codeinfra.Output[Optional[str]]:
        return codeinfra.get(self, "another_value")

    @property
    @codeinfra.getter(name="nestedValue")
    def nested_value(self) -> codeinfra.Output[Optional[NestedNewOutput]]:
        return codeinfra.get(self, "nested_value")

    @property
    @codeinfra.getter(name="listOfNestedValues")
    def list_of_nested_values(self) -> codeinfra.Output[Optional[Sequence[NestedNewOutput]]]:
        return codeinfra.get(self, "list_of_nested_values")

    @property
    @codeinfra.getter
    def tags(self) -> codeinfra.Output[Optional[Mapping[str, str]]]:
        return codeinfra.get(self, "tags")

    @property
    @codeinfra.getter
    def items(self) -> codeinfra.Output[Optional[Mapping[str, str]]]:
        return codeinfra.get(self, "items")

    @property
    @codeinfra.getter
    def keys(self) -> codeinfra.Output[Optional[Sequence[str]]]:
        return codeinfra.get(self, "keys")

    @property
    @codeinfra.getter
    def values(self) -> codeinfra.Output[Optional[Sequence[str]]]:
        return codeinfra.get(self, "values")

    @property
    @codeinfra.getter
    def get(self) -> codeinfra.Output[Optional[str]]:
        return codeinfra.get(self, "get")

    @property
    @codeinfra.getter(name="somethingElse")
    def something_else(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "something_else")

    def translate_output_property(self, prop):
        raise Exception("Unexpected call to translate_output_property")

    def translate_input_property(self, prop):
        raise Exception("Unexpected call to translate_input_property")


class MyMocks(codeinfra.runtime.Mocks):
    def call(self, args: codeinfra.runtime.MockCallArgs):
        return {}
    def new_resource(self, args: codeinfra.runtime.MockResourceArgs):
        if args.name == "o1":
            assert args.inputs == {"serviceName": "hello"}
        elif args.name == "o2":
            assert args.inputs == {"serviceName": "hi", "nestedValue": {"fooBar": "foo bar", "someValue": 42}}
        elif args.name == "o3":
            assert args.inputs == {"serviceName": "bye", "nestedValue": {"fooBar": "bar foo", "someValue": 24}}
        elif args.name == "o4":
            # The "service_name" key in the user-defined tags dict is translated to "serviceName"
            # due to it being in the case tables.
            assert args.inputs == {"serviceName": "sn", "tags": {"serviceName": "my-service"}}
        elif args.name == "o5":
            assert args.inputs == {"serviceName": "o5sn", "listOfNestedValues": [{"fooBar": "f", "someValue": 1}]}
        elif args.name == "o6":
            assert args.inputs == {"serviceName": "o6sn", "listOfNestedValues": [{"fooBar": "b", "someValue": 2}]}
        elif args.name == "n1":
            assert args.inputs == {"serviceName": "hi new", "nestedValue": {"fooBar": "noo nar", "someValue": 100}}
        elif args.name == "n2":
            assert args.inputs == {"serviceName": "hello new", "nestedValue": {"fooBar": "2", "someValue": 3}}
        elif args.name == "n3":
            # service_name correctly isn't translated, because the tags dict is a user-defined dict.
            assert args.inputs == {"serviceName": "sn", "tags": {"service_name": "a-service"}}
        elif args.name == "n4":
            assert args.inputs == {"serviceName": "a", "items": {"foo": "bar"}, "keys": ["foo"], "values": ["bar"], "get": "bar"}
        elif args.name == "d1":
            assert args.inputs == {"serviceName": "b", "items": {"hello": "world"}, "keys": ["hello"], "values": ["world"], "get": "world"}
        return [args.name + '_id', {**args.inputs, "somethingElse": "hehe"}]

codeinfra.runtime.set_mocks(MyMocks())

o1 = OldBehavior("o1", service_name="hello")

o2 = OldBehavior("o2", service_name="hi", nested_value=NestedArgs(foo_bar="foo bar", some_value=42))
def o2checks(v: NestedOldOutput):
    assert v.foo_bar == "foo bar"
    assert v["fooBar"] == "foo bar" # key remains camelCase because it's not in casing table.
    assert v.some_value == 42
    assert v["some_value"] == 42
    return v
codeinfra.export("o2", o2.nested_value.apply(o2checks))

# "foo_bar" isn't in the casing table, so it must be specified as "fooBar", but "some_value" is,
# so it could be snake_case in the dict.
o3 = OldBehavior("o3", service_name="bye", nested_value={"fooBar": "bar foo", "some_value": 24})
def o3checks(v: NestedOldOutput):
    assert v.foo_bar == "bar foo"
    assert v["fooBar"] == "bar foo" # key remains camelCase because it's not in casing table.
    assert v.some_value == 24
    assert v["some_value"] == 24
    return v
codeinfra.export("o3", o3.nested_value.apply(o3checks))

o4 = OldBehavior("o4", service_name="sn", tags={"service_name": "my-service"})
def o4checks(v: Mapping[str, str]):
    assert v == {"service_name": "my-service"}
    return v
codeinfra.export("o4", o4.tags.apply(o4checks))

o5 = OldBehavior("o5", service_name="o5sn", list_of_nested_values=[NestedArgs(foo_bar="f", some_value=1)])
def o5checks(v: Sequence[NestedOldOutput]):
    assert v[0].foo_bar == "f"
    assert v[0]["fooBar"] == "f" # key remains camelCase because it's not in casing table.
    assert v[0].some_value == 1
    assert v[0]["some_value"] == 1
    return v
codeinfra.export("o5", o5.list_of_nested_values.apply(o5checks))

o6 = OldBehavior("o6", service_name="o6sn", list_of_nested_values=[{"fooBar": "b", "some_value": 2}])
def o6checks(v: Sequence[NestedOldOutput]):
    assert v[0].foo_bar == "b"
    assert v[0]["fooBar"] == "b" # key remains camelCase because it's not in casing table.
    assert v[0].some_value == 2
    assert v[0]["some_value"] == 2
    return v
codeinfra.export("o6", o6.list_of_nested_values.apply(o6checks))


n1 = NewBehavior("n1", service_name="hi new", nested_value=NestedArgs(foo_bar="noo nar", some_value=100))
def n1checks(v: NestedNewOutput):
    assert v.foo_bar == "noo nar"
    assert v["foo_bar"] == "noo nar" # key consistently snake_case
    assert v.some_value == 100
    assert v["some_value"] == 100
    return v
codeinfra.export("n1", n1.nested_value.apply(n1checks))

n2 = NewBehavior("n2", service_name="hello new", nested_value={"foo_bar": "2", "some_value": 3})
def n2checks(v: NestedNewOutput):
    assert v.foo_bar == "2"
    assert v["foo_bar"] == "2" # key consistently snake_case
    assert v.some_value == 3
    assert v["some_value"] == 3
    return v
codeinfra.export("n2", n2.nested_value.apply(n2checks))

n3 = NewBehavior("n3", service_name="sn", tags={"service_name": "a-service"})
def n3checks(v: Mapping[str, str]):
    assert v == {"service_name": "a-service"}
    return v
codeinfra.export("n3", n3.tags.apply(n3checks))

n4 = NewBehavior("n4", service_name="a", items={"foo": "bar"}, keys=["foo"], values=["bar"], get="bar")

d1 = NewBehaviorDict("d1", NewBehaviorDictArgs(
    service_name="b", items={"hello": "world"}, keys=["hello"], values=["world"], get="world"))
