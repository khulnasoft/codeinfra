import codeinfra
import codeinfra_random as random

random_password = random.RandomPassword("randomPassword",
    length=16,
    special=True,
    override_special="_%@")
codeinfra.export("password", random_password.result)
