import * as codeinfra from "@codeinfra/codeinfra";
import * as secret from "@codeinfra/secret";

const res = new secret.Resource("res", {
    "private": "closed",
    "public": "open",
    privateData: {
        "private": "closed",
        "public": "open",
    },
    publicData: {
        "private": "closed",
        "public": "open",
    },
});
