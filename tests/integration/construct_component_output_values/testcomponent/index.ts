// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import * as provider from "@codeinfra/codeinfra/provider";

interface BarArgs {
    tags?: codeinfra.Input<{[key: string]: codeinfra.Input<string>}>;
}

interface FooArgs {
    something?: codeinfra.Input<string>;
}

interface ComponentArgs {
    bar?: codeinfra.Input<BarArgs>;
    foo?: FooArgs;
}

class Component extends codeinfra.ComponentResource {
    constructor(name: string, args: ComponentArgs, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, args, opts);

        function isPromise(obj: any): obj is Promise<unknown> {
            return !!obj && obj.then === "function";
        }

        if (!args.foo) {
            throw new Error("expected args.foo to be present");
        }
        if (codeinfra.Output.isInstance(args.foo)) {
            throw new Error("expected args.foo not to be an instance of codeinfra.Output");
        }
        if (!args.foo.something) {
            throw new Error("expected args.foo.something to be present");
        }
        if (args.foo.something !== "hello") {
            throw new Error(`expected args.foo.something to equal "hello" but got "${args.foo.something}"`);
        }

        if (!args.bar) {
            throw new Error("expected args.bar to be present");
        }
        if (codeinfra.Output.isInstance(args.bar)) {
            throw new Error("expected args.bar not to be an instance of codeinfra.Output");
        }
        if (isPromise(args.bar)) {
            throw new Error("expected args.bar not to be a promise");
        }
        if (!args.bar.tags) {
            throw new Error("expected args.bar.tags to be present");
        }
        if (codeinfra.Output.isInstance(args.bar.tags)) {
            throw new Error("expected args.bar.tags not to be an instance of codeinfra.Output");
        }
        if (isPromise(args.bar.tags)) {
            throw new Error("expected args.bar.tags not to be a promise");
        }
        if (args.bar.tags.a !== "world") {
            throw new Error(`expected args.bar.tags.a to equal "world" but got "${args.bar.tags.a}"`);
        }
        if (!codeinfra.Output.isInstance(args.bar.tags.b)) {
            throw new Error(`expected args.bar.tags.b to be an instance of codeinfra.Output`);
        }
        args.bar.tags.b.apply(v => {
            if (v != "shh") {
                throw new Error(`expected args.bar.tags.b to equal "shh" but got "${v}"`);
            }
        });
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    async construct(name: string, type: string, inputs: codeinfra.Inputs,
              options: codeinfra.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

        const component = new Component(name, <ComponentArgs>inputs, options);
        return {
            urn: component.urn,
            state: inputs,
        };
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
