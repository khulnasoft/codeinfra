import * as codeinfra from "@codeinfra/codeinfra";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: codeinfra.output("it's a secret to everybody")
});

export const withSecret = new R("withSecret", {
    prefix: codeinfra.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: codeinfra.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]
});
