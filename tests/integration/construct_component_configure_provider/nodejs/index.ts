// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import * as metaprovider from "@codeinfra/metaprovider";
import * as codeinfra from "@codeinfra/codeinfra";
import * as tls from "@codeinfra/tls";

export = async () => {
  const config = new codeinfra.Config();
  const proxy = config.require("proxy");

  // This resource is just here to generate unknown outputs to test propagating unknowns; there seems to be no way to
  // construct an unknown output in Node.
  const helperPrivateKey = new tls.PrivateKey("helper-private-key", {
    algorithm: "ECDSA",
    ecdsaCurve: "P384",
  });

  const configurer = new metaprovider.Configurer("configurer", {
    tlsProxy: helperPrivateKey.publicKeyPem.apply(_ => proxy) // apply trick makes it unknown at preview
  });

  const key = new tls.PrivateKey("my-private-key", {
    algorithm: "ECDSA",
    ecdsaCurve: "P384",
  }, {
    provider: await configurer.tlsProvider()
  });

  const mix = await configurer.objectMix();

  const key2 = new tls.PrivateKey("my-private-key-e", {
    algorithm: "ECDSA",
    ecdsaCurve: "P384",
  }, {
    provider: mix.provider
  });

  return {
    keyAlgo: key.algorithm,
    keyAlgo2: key2.algorithm,
    meaningOfLife: (await configurer.meaningOfLife() + 1 - 1),
    meaningOfLife2: (await (mix.meaningOfLife||0) + 1 - 1)
  };
};
