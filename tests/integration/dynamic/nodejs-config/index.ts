// Copyright 2016-2023, KhulnaSoft Ltd.

import * as codeinfra from "@codeinfra/codeinfra";

class CustomResource extends codeinfra.dynamic.Resource {
    public readonly authenticated!: codeinfra.Output<string>;
    public readonly color!: codeinfra.Output<string>;

    constructor(name: string, opts?: codeinfra.ResourceOptions) {
        super(
            new DummyResourceProvider(),
            name,
            {
                authenticated: undefined,
                color: undefined
            },
            opts,
            "custom-provider",
            "CustomResource",
        );
    }
}

class DummyResourceProvider implements codeinfra.dynamic.ResourceProvider {
    private password: string;
    private color: string;

    async configure(req: codeinfra.dynamic.ConfigureRequest): Promise<any> {
        this.password = req.config.require("password");
        this.color = req.config.get("colors:banana") ?? "blue";
    }

    async create(props: any): Promise<codeinfra.dynamic.CreateResult> {
        return {
            id: "resource-id",
            outs: {
                authenticated: this.password === "s3cret" ? "200" : "401",
                color: this.color,
            },
        };
    }
}

const resource = new CustomResource("resource-name");

export const authenticated = resource.authenticated;
export const color = resource.color;
