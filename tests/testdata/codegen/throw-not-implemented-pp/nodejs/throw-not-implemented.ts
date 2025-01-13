import * as codeinfra from "@codeinfra/codeinfra";

function notImplemented(message: string) {
    throw new Error(message);
}

export const result = notImplemented("expression here is not implemented yet");
