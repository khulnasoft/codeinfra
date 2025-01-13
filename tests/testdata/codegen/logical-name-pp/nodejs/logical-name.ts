import * as codeinfra from "@codeinfra/codeinfra";
import * as random from "@codeinfra/random";

export = async () => {
    const config = new codeinfra.Config();
    const configLexicalName = config.require("cC-Charlie_charlie.😃⁉️");
    const resourceLexicalName = new random.RandomPet("aA-Alpha_alpha.🤯⁉️", {prefix: configLexicalName});
    return {
        "bB-Beta_beta.💜⁉": resourceLexicalName.id,
        "dD-Delta_delta.🔥⁉": resourceLexicalName.id,
    };
}
