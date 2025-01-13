import * as codeinfra from "@codeinfra/codeinfra";
import * as random from "@codeinfra/random";
import { SimpleComponent } from "./simpleComponent";

interface ExampleComponentArgs {
    /**
     * A simple input
     */
    input: codeinfra.Input<string>,
    /**
     * The main CIDR blocks for the VPC
     * It is a map of strings
     */
    cidrBlocks: codeinfra.Input<Record<string, codeinfra.Input<string>>>,
    /**
     * GitHub app parameters, see your github app. Ensure the key is the base64-encoded `.pem` file (the output of `base64 app.private-key.pem`, not the content of `private-key.pem`).
     */
    githubApp: {
        id?: codeinfra.Input<string>,
        keyBase64?: codeinfra.Input<string>,
        webhookSecret?: codeinfra.Input<string>,
    },
    /**
     * A list of servers
     */
    servers: {
        name?: codeinfra.Input<string>,
    }[],
    /**
     * A map between for zones
     */
    deploymentZones: Record<string, {
        zone?: codeinfra.Input<string>,
    }>,
    ipAddress: codeinfra.Input<number[]>,
}

export class ExampleComponent extends codeinfra.ComponentResource {
    public result: codeinfra.Output<string>;
    constructor(name: string, args: ExampleComponentArgs, opts?: codeinfra.ComponentResourceOptions) {
        super("components:index:ExampleComponent", name, args, opts);
        const password = new random.RandomPassword(`${name}-password`, {
            length: 16,
            special: true,
            overrideSpecial: args.input,
        }, {
            parent: this,
        });

        const githubPassword = new random.RandomPassword(`${name}-githubPassword`, {
            length: 16,
            special: true,
            overrideSpecial: args.githubApp.webhookSecret,
        }, {
            parent: this,
        });

        // Example of iterating a list of objects
        const serverPasswords: random.RandomPassword[] = [];
        for (const range = {value: 0}; range.value < args.servers.length; range.value++) {
            serverPasswords.push(new random.RandomPassword(`${name}-serverPasswords-${range.value}`, {
                length: 16,
                special: true,
                overrideSpecial: args.servers[range.value].name,
            }, {
            parent: this,
        }));
        }

        // Example of iterating a map of objects
        const zonePasswords: random.RandomPassword[] = [];
        for (const range of Object.entries(args.deploymentZones).map(([k, v]) => ({key: k, value: v}))) {
            zonePasswords.push(new random.RandomPassword(`${name}-zonePasswords-${range.key}`, {
                length: 16,
                special: true,
                overrideSpecial: range.value.zone,
            }, {
            parent: this,
        }));
        }

        const simpleComponent = new SimpleComponent(`${name}-simpleComponent`, {
            parent: this,
        });

        this.result = password.result;
        this.registerOutputs({
            result: password.result,
        });
    }
}
