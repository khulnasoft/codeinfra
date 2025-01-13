# Copyright 2016-2020, KhulnaSoft Ltd.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import unittest

from codeinfra._types import resource_types
import codeinfra


class Resource1(codeinfra.Resource):
    pass


class Resource2(codeinfra.Resource):
    foo: codeinfra.Output[str]


class Resource3(codeinfra.Resource):
    nested: codeinfra.Output["Nested"]


class Resource4(codeinfra.Resource):
    nested_value: codeinfra.Output["Nested"] = codeinfra.property("nestedValue")


class Resource5(codeinfra.Resource):
    @property
    @codeinfra.getter
    def foo(self) -> codeinfra.Output[str]: ...  # type: ignore


class Resource6(codeinfra.Resource):
    @property
    @codeinfra.getter
    def nested(self) -> codeinfra.Output["Nested"]: ...  # type: ignore


class Resource7(codeinfra.Resource):
    @property
    @codeinfra.getter(name="nestedValue")
    def nested_value(self) -> codeinfra.Output["Nested"]: ...  # type: ignore


class Resource8(codeinfra.Resource):
    foo: codeinfra.Output


class Resource9(codeinfra.Resource):
    @property
    @codeinfra.getter
    def foo(self) -> codeinfra.Output: ...  # type: ignore


class Resource10(codeinfra.Resource):
    foo: str


class Resource11(codeinfra.Resource):
    @property
    @codeinfra.getter
    def foo(self) -> str: ...  # type: ignore


class Resource12(codeinfra.Resource):
    @property
    @codeinfra.getter
    def foo(self): ...  # type: ignore


@codeinfra.output_type
class Nested:
    first: str
    second: str


class ResourceTypesTests(unittest.TestCase):
    def test_resource_types(self):
        self.assertEqual({}, resource_types(Resource1))

        self.assertEqual({"foo": str}, resource_types(Resource2))
        self.assertEqual({"nested": Nested}, resource_types(Resource3))
        self.assertEqual({"nestedValue": Nested}, resource_types(Resource4))

        self.assertEqual({"foo": str}, resource_types(Resource5))
        self.assertEqual({"nested": Nested}, resource_types(Resource6))
        self.assertEqual({"nestedValue": Nested}, resource_types(Resource7))

        # Non-generic Output excluded from types.
        self.assertEqual({}, resource_types(Resource8))
        self.assertEqual({}, resource_types(Resource9))

        # Type annotations not using Output.
        self.assertEqual({"foo": str}, resource_types(Resource10))
        self.assertEqual({"foo": str}, resource_types(Resource11))

        # No return type annotation from the property getter.
        self.assertEqual({}, resource_types(Resource12))
