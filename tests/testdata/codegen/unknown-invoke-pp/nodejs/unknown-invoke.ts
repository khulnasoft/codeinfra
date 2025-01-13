import * as codeinfra from "@codeinfra/codeinfra";
import * as unknown from "@codeinfra/unknown";

const data = unknown.index.getData({
    input: "hello",
});
const values = unknown.eks.moduleValues({});
export const content = data.content;
