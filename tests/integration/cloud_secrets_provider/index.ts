import * as codeinfra from "@codeinfra/codeinfra";

const config = new codeinfra.Config();

export const out = config.requireSecret("mysecret");
