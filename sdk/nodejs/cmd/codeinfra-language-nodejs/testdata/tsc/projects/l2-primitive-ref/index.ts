import * as codeinfra from "@codeinfra/codeinfra";
import * as primitive_ref from "@codeinfra/primitive-ref";

const res = new primitive_ref.Resource("res", {data: {
    boolean: false,
    float: 2.17,
    integer: -12,
    string: "Goodbye",
    boolArray: [
        false,
        true,
    ],
    stringMap: {
        two: "turtle doves",
        three: "french hens",
    },
}});
