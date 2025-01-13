import * as codeinfra from "@codeinfra/codeinfra";
import { First } from "./first";
import { Second } from "./second";

const [secondPasswordLength, resolveSecondPasswordLength] = codeinfra.deferredOutput<number>();
const first = new First("first", {passwordLength: secondPasswordLength});
const second = new Second("second", {petName: first.petName});
resolveSecondPasswordLength(second.passwordLength);
const [loopingOverMany, resolveLoopingOverMany] = codeinfra.deferredOutput<Array<number>>();
const another = new First("another", {passwordLength: loopingOverMany.apply(loopingOverMany => loopingOverMany.length)});
const many: Second[] = [];
for (const range = {value: 0}; range.value < 10; range.value++) {
    many.push(new Second(`many-${range.value}`, {petName: another.petName}));
}
resolveLoopingOverMany(codeinfra.output(many.map((v, k) => [k, v]).map(([_, v]) => (v.passwordLength))));
