import * as codeinfra from "@codeinfra/codeinfra";
import * as provider from "@codeinfra/codeinfra/provider";

interface ComponentArgs {
    message: codeinfra.Input<string>;
    nested: codeinfra.Input<{
        value: codeinfra.Input<string>;
    }>;
}

class Component extends codeinfra.ComponentResource {
    constructor(name: string, args: ComponentArgs, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, {}, opts);

        // These `apply`s should not run.
        codeinfra.output(args.message).apply(v => { console.log("should not run (message)"); process.exit(1); });
        codeinfra.output(args.nested).apply(v => { console.log("should not run (nested)"); process.exit(1); });
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    construct(name: string, type: string, inputs: codeinfra.Inputs,
              options: codeinfra.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

        const component = new Component(name, <ComponentArgs>inputs, options);
        return Promise.resolve({
            urn: component.urn,
            state: {},
        });
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
