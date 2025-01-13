// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class Provider extends codeinfra.ProviderResource {
    public readonly message!: codeinfra.Output<string>;

    constructor(name: string, message: string, opts?: codeinfra.ResourceOptions) {
        super("testcomponent", name, { message }, opts);
    }
}

class Component extends codeinfra.ComponentResource {
    public readonly message!: codeinfra.Output<string>;

    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        const inputs = {
            message: undefined /*out*/,
        };
        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

class LocalComponent extends codeinfra.ComponentResource {
    public readonly message: codeinfra.Output<string>;

    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("my:index:LocalComponent", name, {}, opts);

        const component = new Component(`${name}-mycomponent`, { parent: this });
        this.message = component.message;
    }
}

const provider = new Provider("myprovider", "hello world")
const component = new Component("mycomponent", { provider });
const localComponent = new LocalComponent("mylocalcomponent", { providers: [provider] });

export const message = component.message;
export const nestedMessage = localComponent.message;
