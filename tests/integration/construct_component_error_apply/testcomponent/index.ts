// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import * as provider from "@codeinfra/codeinfra/provider";

class Component extends codeinfra.ComponentResource {
    public readonly foo: codeinfra.Output<string>;

    constructor(name: string, foo: codeinfra.Input<string>, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, {}, opts);

        this.foo = codeinfra.output(foo);

        this.registerOutputs({
            foo: this.foo,
        })
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    construct(name: string, type: string, inputs: codeinfra.Inputs,
              options: codeinfra.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

       const foo = codeinfra.output("").apply(a => {
           throw new Error("intentional error from within an apply");
           return a;
       });


        const component = new Component(name, foo);
        return Promise.resolve({
            urn: component.urn,
            state: {
                foo: component.foo
            },
        });
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
