// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class Provider implements codeinfra.dynamic.ResourceProvider {
    public readonly create: (inputs: any) => Promise<codeinfra.dynamic.CreateResult>;

    constructor(num: number) {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: { value: num }
            }
        }
    }
}


class FirstResource extends codeinfra.dynamic.Resource {
    public readonly value: codeinfra.Output<number>;

    private static provider: Provider = new Provider(42);
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}

class SecondResource extends codeinfra.dynamic.Resource {
    public readonly dep: codeinfra.Output<number>;

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: codeinfra.Input<number>) {
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}

const first = new FirstResource("first");
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});