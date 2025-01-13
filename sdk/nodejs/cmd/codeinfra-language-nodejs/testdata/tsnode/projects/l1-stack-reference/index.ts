import * as codeinfra from "@codeinfra/codeinfra";

const ref = new codeinfra.StackReference("ref", {name: "organization/other/dev"});
export const plain = ref.getOutput("plain");
export const secret = ref.getOutput("secret");
