# Copyright 2016-2021, KhulnaSoft Ltd.
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

from typing import Optional

from codeinfra._types import input_type_types
import codeinfra


@codeinfra.input_type
class Foo:
    @property
    @codeinfra.getter()
    def bar(self) -> codeinfra.Input[str]: ...  # type: ignore


@codeinfra.input_type
class MySimpleInputType:
    a: str
    b: Optional[str]
    c: codeinfra.Input[str]
    d: Optional[codeinfra.Input[str]]
    e: Foo
    f: Optional[Foo]
    g: codeinfra.Input[Foo]
    h: Optional[codeinfra.Input[Foo]]
    i: codeinfra.InputType[Foo]
    j: Optional[codeinfra.InputType[Foo]]
    k: codeinfra.Input[codeinfra.InputType[Foo]]
    l: Optional[codeinfra.Input[codeinfra.InputType[Foo]]]


@codeinfra.input_type
class MyPropertiesInputType:
    @property
    @codeinfra.getter()
    def a(self) -> str: ...  # type: ignore

    @property
    @codeinfra.getter()
    def b(self) -> Optional[str]: ...  # type: ignore

    @property
    @codeinfra.getter()
    def c(self) -> codeinfra.Input[str]: ...  # type: ignore

    @property
    @codeinfra.getter()
    def d(self) -> Optional[codeinfra.Input[str]]: ...  # type: ignore

    @property
    @codeinfra.getter()
    def e(self) -> Foo: ...  # type: ignore

    @property
    @codeinfra.getter()
    def f(self) -> Optional[Foo]: ...  # type: ignore

    @property
    @codeinfra.getter()
    def g(self) -> codeinfra.Input[Foo]: ...  # type: ignore

    @property
    @codeinfra.getter()
    def h(self) -> Optional[codeinfra.Input[Foo]]: ...  # type: ignore

    @property
    @codeinfra.getter()
    def i(self) -> codeinfra.InputType[Foo]: ...  # type: ignore

    @property
    @codeinfra.getter()
    def j(self) -> Optional[codeinfra.InputType[Foo]]: ...  # type: ignore

    @property
    @codeinfra.getter()
    def k(self) -> codeinfra.Input[codeinfra.InputType[Foo]]: ...  # type: ignore

    @property
    @codeinfra.getter()
    def l(self) -> Optional[codeinfra.Input[codeinfra.InputType[Foo]]]: ...  # type: ignore


class InputTypeTypesTests(unittest.TestCase):
    def test_input_type_types(self):
        expected = {
            "a": str,
            "b": str,
            "c": str,
            "d": str,
            "e": Foo,
            "f": Foo,
            "g": Foo,
            "h": Foo,
            "i": Foo,
            "j": Foo,
            "k": Foo,
            "l": Foo,
        }
        self.assertEqual(expected, input_type_types(MySimpleInputType))
        self.assertEqual(expected, input_type_types(MyPropertiesInputType))
