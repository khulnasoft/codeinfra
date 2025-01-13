import * as codeinfra from "@codeinfra/codeinfra";
import * as large from "@codeinfra/large";

const res = new large.String("res", {value: "hello world"});
export const output = res.value;
