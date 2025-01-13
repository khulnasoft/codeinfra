import * as codeinfra from "@codeinfra/codeinfra";
import * as random from "@codeinfra/random";

export class AnotherComponent extends codeinfra.ComponentResource {
    constructor(name: string, opts?: codeinfra.ComponentResourceOptions) {
        super("components:index:AnotherComponent", name, {}, opts);
        const firstPassword = new random.RandomPassword(`${name}-firstPassword`, {
            length: 16,
            special: true,
        }, {
            parent: this,
        });

        this.registerOutputs();
    }
}
