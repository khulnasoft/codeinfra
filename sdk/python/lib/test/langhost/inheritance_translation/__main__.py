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

import codeinfra


class MyResource(codeinfra.CustomResource):
    def __init__(self, name):
        # Pass a @codeinfra.input_type to opt-in to new translation behavior.
        @codeinfra.input_type
        class Args:
            pass

        props = Args()
        props.__dict__["some_value"] = None
        props.__dict__["another_value"] = None
        super().__init__("test:index:MyResource", name, props)

    @property
    @codeinfra.getter(name="someValue")
    def some_value(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "some_value")

    @property
    @codeinfra.getter(name="anotherValue")
    def another_value(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "another_value")


class MyResourceSubclass(MyResource):
    combined_values: codeinfra.Output[str]

    def __init__(self, name):
        super().__init__(name)
        self.combined_values = codeinfra.Output.concat(
            self.some_value, " ", self.another_value
        )


class MyResourceSubclassSubclass(MyResourceSubclass):
    new_value: codeinfra.Output[str]

    def __init__(self, name):
        super().__init__(name)
        self.new_value = codeinfra.Output.concat(self.combined_values, "!")


class MyLegacyTranslationResource(codeinfra.CustomResource):
    def __init__(self, name):
        # Pass a regular dict to use the old translation behavior.
        props = {}
        props["some_value"] = None
        props["another_value"] = None
        super().__init__("test:index:MyResource", name, props)

    def translate_output_property(self, prop: str) -> str:
        return {
            "someValue": "some_value",
            "anotherValue": "another_value",
        }.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return {
            "some_value": "someValue",
            "another_value": "anotherValue",
        }.get(prop) or prop

    @property
    @codeinfra.getter(name="someValue")
    def some_value(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "some_value")

    @property
    @codeinfra.getter(name="anotherValue")
    def another_value(self) -> codeinfra.Output[str]:
        return codeinfra.get(self, "another_value")


class MyLegacyTranslationResourceSubclass(MyLegacyTranslationResource):
    combined_values: codeinfra.Output[str]

    def __init__(self, name):
        super().__init__(name)
        self.combined_values = codeinfra.Output.concat(
            self.some_value, " ", self.another_value
        )


class MyLegacyTranslationResourceSubclassSubclass(MyLegacyTranslationResourceSubclass):
    new_value: codeinfra.Output[str]

    def __init__(self, name):
        super().__init__(name)
        self.new_value = codeinfra.Output.concat(self.combined_values, "!")


r1 = MyResourceSubclass("testResource")
r2 = MyResourceSubclassSubclass("testResource")
codeinfra.export("r1.some_value", r1.some_value)
codeinfra.export("r1.another_value", r1.another_value)
codeinfra.export("r1.combined_values", r1.combined_values)
codeinfra.export("r2.some_value", r2.some_value)
codeinfra.export("r2.another_value", r2.another_value)
codeinfra.export("r2.combined_values", r2.combined_values)
codeinfra.export("r2.new_value", r2.new_value)


r3 = MyLegacyTranslationResourceSubclass("testResource")
r4 = MyLegacyTranslationResourceSubclassSubclass("testResource")
codeinfra.export("r3.some_value", r3.some_value)
codeinfra.export("r3.another_value", r3.another_value)
codeinfra.export("r3.combined_values", r3.combined_values)
codeinfra.export("r4.some_value", r4.some_value)
codeinfra.export("r4.another_value", r4.another_value)
codeinfra.export("r4.combined_values", r4.combined_values)
codeinfra.export("r4.new_value", r4.new_value)
