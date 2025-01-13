import codeinfra
from codeinfra import Input
from typing import Optional, Dict, TypedDict, Any
import codeinfra_random as random

class AnotherComponent(codeinfra.ComponentResource):
    def __init__(self, name: str, opts: Optional[codeinfra.ResourceOptions] = None):
        super().__init__("components:index:AnotherComponent", name, {}, opts)

        first_password = random.RandomPassword(f"{name}-firstPassword",
            length=16,
            special=True,
            opts = codeinfra.ResourceOptions(parent=self))

        self.register_outputs()
