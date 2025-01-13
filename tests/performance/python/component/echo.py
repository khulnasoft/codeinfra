# Copyright 2024, KhulnaSoft Ltd.

from typing import Any, Optional

import codeinfra

class Echo(codeinfra.CustomResource):
    def __init__(
        self,
        resource_name: str,
        echo: codeinfra.Input[Any],
        opts: Optional[codeinfra.ResourceOptions] = None,
    ):
        props = {
            "echo": echo,
        }
        super().__init__("testprovider:index:Echo", resource_name, props, opts)

    @property
    @codeinfra.getter
    def echo(self) -> codeinfra.Output[Any]:
        return codeinfra.get(self, "echo")
