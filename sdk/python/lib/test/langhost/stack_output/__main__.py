# Copyright 2016-2018, KhulnaSoft Ltd.
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
from typing import Any, Dict
import codeinfra


class TestClass:
    def __init__(self):
        self.num = 1
        self._private = 2


my_test_class_instance = TestClass()

recursive: Dict[str, Any] = {"a": 1}
recursive["b"] = 2
recursive["c"] = recursive

codeinfra.export("string", "codeinfra")
codeinfra.export("number", 1)
codeinfra.export("boolean", True)
codeinfra.export("list", [])
codeinfra.export("list_with_none", [None])
codeinfra.export("list_of_lists", [[], []])
codeinfra.export(
    "list_of_outputs", [[codeinfra.Output.from_input(1)], codeinfra.Output.from_input([2])]
)
codeinfra.export("set", set(["val"]))
codeinfra.export("dict", {"a": 1})
codeinfra.export("output", codeinfra.Output.from_input(1))
codeinfra.export("class", TestClass())
codeinfra.export("recursive", recursive)
codeinfra.export("duplicate_output_0", my_test_class_instance)
codeinfra.export("duplicate_output_1", my_test_class_instance)
