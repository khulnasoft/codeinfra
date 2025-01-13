import codeinfra
import codeinfra_random as random

foo = random.RandomPet("foo", opts = codeinfra.ResourceOptions(retain_on_delete=True))
