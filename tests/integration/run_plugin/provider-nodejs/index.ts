// Copyright 2024 KhulnaSoft Ltd.

import * as codeinfra from "@codeinfra/codeinfra";
import * as provider from "@codeinfra/codeinfra/provider";

class Provider implements provider.Provider {
    constructor(readonly version: string) {}

    async construct(name: string, type: string, inputs: codeinfra.Inputs,
        options: codeinfra.ComponentResourceOptions): Promise<provider.ConstructResult> {
        const typeName = type.split(":", 3)[2];
        return {
            urn: codeinfra.createUrn(type, name),
            state: {
                "ITS_ALIVE": "IT'S ALIVE!",
            }
        }
    }
}

const prov = new Provider("0.0.1");
codeinfra.provider.main(prov, process.argv.slice(2));
