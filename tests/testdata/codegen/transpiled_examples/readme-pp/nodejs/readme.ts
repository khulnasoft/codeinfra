import * as codeinfra from "@codeinfra/codeinfra";
import * as fs from "fs";

export const strVar = "foo";
export const arrVar = [
    "fizz",
    "buzz",
];
export const readme = fs.readFileSync("./Codeinfra.README.md", "utf8");
