// Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class MyResource extends codeinfra.dynamic.Resource {
    constructor(name: string, props: codeinfra.Inputs, opts?: codeinfra.CustomResourceOptions) {
        super({
            create: async (inputs: any) => {
                return {
                    id: "0",
                    outs: inputs,
                }
            },
        }, name, props, opts);
    }
}

class GetResource extends codeinfra.Resource {
    foo: codeinfra.Output<string>;
    bar: codeinfra.Output<string>;

    constructor(urn: codeinfra.URN) {
        const props = {
            foo: undefined,
            bar: undefined,
        };
        super("unused:unused:unused", "unused", true, props, { urn });
    }
}

const a = new MyResource("a", {
    foo: "foo",
    bar: codeinfra.secret("my-$ecret"),
});

const getFoo = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.foo
});

const getBar = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.bar
});


export const foo = getFoo;
export const secret = getBar;