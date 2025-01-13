# Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

from typing import Optional

import codeinfra

class Component(codeinfra.ComponentResource):
    """
    Python-level remote component for the component resource
    defined in sibling testcomponent-go directory.
    """

    result: codeinfra.Output[str]

    def __init__(self,
                 name: str,
                 opts: Optional[codeinfra.ResourceOptions] = None):
        props = {"result": None}
        super().__init__("testcomponent:index:Component", name, props, opts, remote=True)


class RandomProvider(codeinfra.ProviderResource):
    """
    Provider for the testprovider:index:Random resource.

    Implemented in tests/testprovider.
    """

    def __init__(self, name, opts: Optional[codeinfra.ResourceOptions]=None):
        super().__init__("testprovider", name, {}, opts)


explicit_provider = RandomProvider("explicit")

# Should pick up the default provider.
Component("uses_default")

# Should use the provider passed in as an argument.
Component("uses_provider", opts=codeinfra.ResourceOptions(
    provider=explicit_provider,
))

# Should use the provider passed in as an argument
Component("uses_providers", opts=codeinfra.ResourceOptions(
    providers=[explicit_provider],
))

# Should use the provider passed in as an argument
Component("uses_providers_map", opts=codeinfra.ResourceOptions(
    providers={"testprovider": explicit_provider},
))
