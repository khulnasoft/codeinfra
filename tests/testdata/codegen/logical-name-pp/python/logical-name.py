import codeinfra
import codeinfra_random as random

config = codeinfra.Config()
config_lexical_name = config.require("cC-Charlie_charlie.😃⁉️")
resource_lexical_name = random.RandomPet("aA-Alpha_alpha.🤯⁉️", prefix=config_lexical_name)
codeinfra.export("bB-Beta_beta.💜⁉", resource_lexical_name.id)
codeinfra.export("dD-Delta_delta.🔥⁉", resource_lexical_name.id)
