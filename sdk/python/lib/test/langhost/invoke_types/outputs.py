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

import codeinfra

import outputs


@codeinfra.output_type
class MyFunctionNestedResult:
    first_value: str = codeinfra.property("firstValue")
    second_value: float = codeinfra.property("secondValue")


@codeinfra.output_type
class MyFunctionResult:
    # Deliberately using a qualified (with `outputs.`) forward reference
    # to mimic our provider codegen, to ensure the type can be resolved.
    nested: "outputs.MyFunctionNestedResult"


@codeinfra.output_type
class MyOtherFunctionNestedResult:
    def __init__(self, first_value: str, second_value: float):
        codeinfra.set(self, "first_value", first_value)
        codeinfra.set(self, "second_value", second_value)

    @property
    @codeinfra.getter(name="firstValue")
    def first_value(self) -> str: ...  # type: ignore

    @property
    @codeinfra.getter(name="secondValue")
    def second_value(self) -> float: ...  # type: ignore


@codeinfra.output_type
class MyOtherFunctionResult:
    def __init__(self, nested: "outputs.MyOtherFunctionNestedResult"):
        codeinfra.set(self, "nested", nested)

    @property
    @codeinfra.getter
    # Deliberately using a qualified (with `outputs.`) forward reference
    # to mimic our provider codegen, to ensure the type can be resolved.
    def nested(self) -> "outputs.MyOtherFunctionNestedResult": ...
