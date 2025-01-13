import * as codeinfra from "@codeinfra/codeinfra";
import * as random from "@codeinfra/random";

const randomPassword = new random.RandomPassword("randomPassword", {
    length: 16,
    special: true,
    overrideSpecial: "_%@",
});
export const password = randomPassword.result;
