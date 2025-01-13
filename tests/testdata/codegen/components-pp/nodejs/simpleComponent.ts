import * as codeinfra from "@codeinfra/codeinfra";
import * as random from "@codeinfra/random";

export class SimpleComponent extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("components:index:SimpleComponent", name, {}, opts);
        const firstPassword = new random.RandomPassword(`${name}-firstPassword`, {
            length: 16,
            special: true,
        }, {
            parent: this,
        });

        const secondPassword = new random.RandomPassword(`${name}-secondPassword`, {
            length: 16,
            special: true,
        }, {
            parent: this,
        });

        this.registerOutputs();
    }
}
