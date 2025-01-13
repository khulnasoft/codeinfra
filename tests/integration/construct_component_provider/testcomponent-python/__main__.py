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

from typing import Optional
import sys

from semver import VersionInfo as SemverVersion

from codeinfra import Inputs, ComponentResource, ResourceOptions
import codeinfra
import codeinfra.provider as provider


VERSION = "0.0.1"


class Provider(codeinfra.ProviderResource):
    message: codeinfra.Output[str]

    def __init__(self, name: str, opts: Optional[codeinfra.ResourceOptions] = None) -> None:
        props = {
            "message": None, # out
        }
        super().__init__("testcomponent", name, props, opts)


class Component(ComponentResource):
    message: codeinfra.Output[str]

    def __init__(self, name: str, opts: Optional[ResourceOptions] = None):
        super().__init__("testcomponent:index:Component", name, {}, opts)
        provider = self.get_provider("testcomponent::")
        if not isinstance(provider, Provider):
            raise Exception(f"provider is not an instance of Provider: {provider}")
        self.message = provider.message
        self.register_outputs({
            "message": self.message,
        })


class Package(codeinfra.runtime.ResourcePackage):
    _version = SemverVersion.parse(VERSION)

    def version(self):
        return self._version

    def construct_provider(self, name: str, typ: str, urn: str) -> codeinfra.ProviderResource:
        if typ != "codeinfra:providers:testcomponent":
            raise Exception(f"unknown provider type {typ}")
        return Provider(name, codeinfra.ResourceOptions(urn=urn))


class ProviderServer(provider.Provider):
    def __init__(self):
        super().__init__(VERSION)
        codeinfra.runtime.register_resource_package("testcomponent", Package())

    def construct(self, name: str, resource_type: str, inputs: Inputs,
                  options: Optional[ResourceOptions] = None) -> provider.ConstructResult:

        if resource_type != "testcomponent:index:Component":
            raise Exception(f"unknown resource type {resource_type}")

        component = Component(name, options)

        return provider.ConstructResult(
            urn=component.urn,
            state={
                "message": component.message,
            })


if __name__ == "__main__":
    provider.main(ProviderServer(), sys.argv[1:])
