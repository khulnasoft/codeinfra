import * as codeinfra from "@codeinfra/codeinfra";

export function willThrow() {
    if (true) {
        codeinfra.log.error("Oh no!");
        throw new Error("this is a test error");
    }
}

willThrow();
