import * as codeinfra from "@codeinfra/codeinfra";
import * as dynamic from "@codeinfra/codeinfra/dynamic";

export interface RArgs {
    prefix: codeinfra.Input<string>
}

const provider: codeinfra.dynamic.ResourceProvider = {
    async create(inputs) {
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};
    }
}

export class R extends dynamic.Resource {
    public prefix!: codeinfra.Output<string>;

    constructor(name: string, props: RArgs, opts?: codeinfra.CustomResourceOptions) {
        super(provider, name, props, opts)
    }
}