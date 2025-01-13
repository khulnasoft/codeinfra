import * as codeinfra from "@codeinfra/codeinfra";
import * as unknown from "@codeinfra/unknown";

const provider = new codeinfra.providers.Unknown("provider", {});
const main = new unknown.index.Main("main", {
    first: "hello",
    second: {
        foo: "bar",
    },
});
const fromModule: unknown.eks.Example[] = [];
for (const range = {value: 0}; range.value < 10; range.value++) {
    fromModule.push(new unknown.eks.Example(`fromModule-${range.value}`, {associatedMain: main.id}));
}
export const mainId = main.id;
export const values = fromModule.values.first;
