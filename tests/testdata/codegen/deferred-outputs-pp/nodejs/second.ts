import * as codeinfra from "@codeinfra/codeinfra";
import * as random from "@codeinfra/random";

interface SecondArgs {
    petName: codeinfra.Input<string>,
}

export class Second extends codeinfra.ComponentResource {
    public passwordLength: codeinfra.Output<number>;
    constructor(name: string, args: SecondArgs, opts?: codeinfra.ComponentResourceOptions) {
        super("components:index:Second", name, args, opts);
        const randomPet = new random.RandomPet(`${name}-randomPet`, {length: args.petName.length}, {
            parent: this,
        });

        const password = new random.RandomPassword(`${name}-password`, {
            length: 16,
            special: true,
            numeric: false,
        }, {
            parent: this,
        });

        this.passwordLength = password.length;
        this.registerOutputs({
            passwordLength: password.length,
        });
    }
}
