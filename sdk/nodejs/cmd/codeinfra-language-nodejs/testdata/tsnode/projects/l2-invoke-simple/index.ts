import * as codeinfra from "@codeinfra/codeinfra";
import * as simple_invoke from "@codeinfra/simple-invoke";

export const hello = simple_invoke.myInvokeOutput({
    value: "hello",
}).apply(invoke => invoke.result);
export const goodbye = simple_invoke.myInvokeOutput({
    value: "goodbye",
}).apply(invoke => invoke.result);
