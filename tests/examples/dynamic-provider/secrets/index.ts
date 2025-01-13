// Copyright 2023, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import * as dynamic from "@codeinfra/codeinfra/dynamic";

const config = new codeinfra.Config();
const password = config.requireSecret("password");

class SimpleProvider implements codeinfra.dynamic.ResourceProvider {
    async create(inputs: any) {
        // Need to use `password.get()` to get the underlying value of the secret from within the serialzied code.  
        // This simulates using this as a credential to talk to an external system.
        return { id: "0", outs: { authenticated: password.get() == "s3cret" ? "200" : "401" } };
    }
}

class SimpleResource extends dynamic.Resource {
    authenticated!: codeinfra.Output<string>;
    constructor(name: string) {
        super(new SimpleProvider(), name, { authenticated: undefined }, undefined);
    }
}

let r = new SimpleResource("foo");
export const out = r.authenticated;

