import * as codeinfra from "@codeinfra/codeinfra";
import * as simple from "@codeinfra/simple";
import * as simple_invoke from "@codeinfra/simple-invoke";

const res = new simple.Resource("res", {value: true});
export const nonSecret = simple_invoke.secretInvokeOutput({
    value: "hello",
    secretResponse: false,
}).apply(invoke => invoke.response);
export const firstSecret = simple_invoke.secretInvokeOutput({
    value: "hello",
    secretResponse: res.value,
}).apply(invoke => invoke.response);
export const secondSecret = simple_invoke.secretInvokeOutput({
    value: codeinfra.secret("goodbye"),
    secretResponse: false,
}).apply(invoke => invoke.response);
