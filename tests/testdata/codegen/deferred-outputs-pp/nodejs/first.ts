import * as codeinfra from "@codeinfra/codeinfra";
import * as random from "@codeinfra/random";

interface FirstArgs {
    passwordLength: codeinfra.Input<number>,
}

export class First extends codeinfra.ComponentResource {
    public petName: codeinfra.Output<string>;
    public password: codeinfra.Output<string>;
    constructor(name: string, args: FirstArgs, opts?: codeinfra.ComponentResourceOptions) {
        super("components:index:First", name, args, opts);
        const randomPet = new random.RandomPet(`${name}-randomPet`, {}, {
            parent: this,
        });

        const randomPassword = new random.RandomPassword(`${name}-randomPassword`, {length: args.passwordLength}, {
            parent: this,
        });

        this.petName = randomPet.id;
        this.password = randomPassword.result;
        this.registerOutputs({
            petName: randomPet.id,
            password: randomPassword.result,
        });
    }
}
