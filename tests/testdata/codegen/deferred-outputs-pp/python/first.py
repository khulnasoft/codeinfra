import codeinfra
from codeinfra import Input
from typing import Optional, Dict, TypedDict, Any
import codeinfra_random as random

class FirstArgs(TypedDict, total=False):
    passwordLength: Input[float]

class First(codeinfra.ComponentResource):
    def __init__(self, name: str, args: FirstArgs, opts:Optional[codeinfra.ResourceOptions] = None):
        super().__init__("components:index:First", name, args, opts)

        random_pet = random.RandomPet(f"{name}-randomPet", opts = codeinfra.ResourceOptions(parent=self))

        random_password = random.RandomPassword(f"{name}-randomPassword", length=args["passwordLength"],
        opts = codeinfra.ResourceOptions(parent=self))

        self.petName = random_pet.id
        self.password = random_password.result
        self.register_outputs({
            'petName': random_pet.id, 
            'password': random_password.result
        })