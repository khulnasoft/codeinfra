import * as codeinfra from "@codeinfra/codeinfra";

const data = [
    1,
    2,
    3,
].map((v, k) => ({key: k, value: v})).map(entry => ({
    usingKey: entry.key,
    usingValue: entry.value,
}));
