// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

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

const component = new Component("mycomponent", {
    providers: {
        "testcomponent": new Provider("myprovider", "hello world"),
    },
});

export const message = component.message;
