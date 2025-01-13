import * as codeinfra from "@codeinfra/codeinfra";
import * as dynamic from "@codeinfra/codeinfra/dynamic";
import * as provider from "@codeinfra/codeinfra/provider";

let currentID = 0;

class Resource extends dynamic.Resource {
    constructor(name: string, echo: codeinfra.Input<any>, opts?: codeinfra.CustomResourceOptions) {
        const provider = {
            create: async (inputs: any) => ({
                id: (currentID++).toString(),
                outs: undefined,
            }),
        };

        super(provider, name, {echo}, opts);
    }
}

class Component extends codeinfra.ComponentResource {
    public readonly echo: codeinfra.Output<any>;
    public readonly childId: codeinfra.Output<codeinfra.ID>;
    public readonly secret: codeinfra.Output<string>;

    constructor(name: string, echo: codeinfra.Input<any>, secret: codeinfra.Output<string>, opts?: codeinfra.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, {}, opts);

        this.echo = codeinfra.output(echo);
        this.childId = (new Resource(`child-${name}`, echo, {parent: this})).id;
        this.secret = secret;

        this.registerOutputs({
            echo: this.echo,
            childId: this.childId,
            secret: this.secret,
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

        const config = new codeinfra.Config();
        const secretKey = "secret";
        const fullSecretKey = `${config.name}:${secretKey}`;
        // use internal codeinfra prop to check secretness
        const isSecret = (codeinfra.runtime as any).isConfigSecret(fullSecretKey); 
        if (!isSecret) {
            throw new Error(`expected config with key "${secretKey}" to be secret.`)
        }
        const secret = config.requireSecret(secretKey);


        const component = new Component(name, inputs["echo"], secret, options);
        return Promise.resolve({
            urn: component.urn,
            state: {
                echo: component.echo,
                childId: component.childId,
                secret: secret,
            },
        });
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
