import * as codeinfra from "@codeinfra/codeinfra";

function singleOrNone<T>(elements: codeinfra.Input<T>[]): codeinfra.Input<T> {
    if (elements.length != 1) {
        throw new Error("singleOrNone expected input list to have a single element");
    }
    return elements[0];
}

export const result = singleOrNone([1]);
