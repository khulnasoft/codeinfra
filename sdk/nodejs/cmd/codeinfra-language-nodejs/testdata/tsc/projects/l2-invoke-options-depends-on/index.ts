import * as codeinfra from "@codeinfra/codeinfra";
import * as simple_invoke from "@codeinfra/simple-invoke";

const explicitProvider = new simple_invoke.Provider("explicitProvider", {});
const first = new simple_invoke.StringResource("first", {text: "first hello"});
const data = simple_invoke.myInvokeOutput({
    value: "hello",
}, {
    dependsOn: [first],
});
const second = new simple_invoke.StringResource("second", {text: data.result});
export const hello = data.result;
