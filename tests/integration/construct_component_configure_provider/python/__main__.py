# Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import codeinfra
import codeinfra_metaprovider
import codeinfra_tls
import helpers

config = codeinfra.Config()
proxy = config.require('proxy')

configurer = codeinfra_metaprovider.Configurer("configurer", tls_proxy=helpers.unknownIfDryRun(proxy))
tls_provider = configurer.tls_provider()

key = codeinfra_tls.PrivateKey("my-private-key",
                            algorithm="ECDSA",
                            ecdsa_curve="P384",
                            opts=codeinfra.ResourceOptions(provider=tls_provider))

codeinfra.export("keyAlgo", key.algorithm)
codeinfra.export("meaningOfLife", configurer.meaning_of_life() + 1 - 1)

mix = configurer.object_mix()
key2 = codeinfra_tls.PrivateKey("my-private-key-2",
                             algorithm="ECDSA",
                             ecdsa_curve="P384",
                             opts=codeinfra.ResourceOptions(provider=mix.provider))

codeinfra.export("keyAlgo2", key2.algorithm)
codeinfra.export("meaningOfLife2", mix.meaning_of_life + 1 - 1)
