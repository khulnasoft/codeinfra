// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class Component extends codeinfra.ComponentResource {
    public readonly result!: codeinfra.Output<string>;

    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        const inputs = {
            result: undefined /*out*/,
        };
        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

class RandomProvider extends codeinfra.ProviderResource {
    constructor(name: string, opts?: codeinfra.ResourceOptions) {
        super("testprovider", name, {}, opts);
    }
}

const explicitProvider = new RandomProvider("explicit");

new Component("uses_default");
new Component("uses_provider", {provider: explicitProvider});
new Component("uses_providers", {providers: [explicitProvider]});
new Component("uses_providers_map", {providers: {testprovider: explicitProvider}});
