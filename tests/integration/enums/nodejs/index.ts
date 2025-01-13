// Copyright 2016-2020, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class PlantProvider implements codeinfra.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<codeinfra.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: inputs,
            };
        };
    }
}

interface RubberTreeArgs {
    readonly farm?: codeinfra.Input<Farm | string>;
    readonly type: codeinfra.Input<RubberTreeVariety>;
}

class RubberTree extends codeinfra.dynamic.Resource {
    public readonly farm!: codeinfra.Output<Farm | string | undefined>;
    public readonly type!: codeinfra.Output<RubberTreeVariety>;

    constructor(name: string, args: RubberTreeArgs) {
        const inputs: codeinfra.Inputs = {
            farm: args.farm,
            type: args.type,
        };
        super(new PlantProvider(), name, inputs, undefined);
    }
}

const Farm = {
    Codeinfra_Planters_Inc_: "Codeinfra Planters Inc.",
    Plants_R_Us: "Plants'R'Us",
} as const;

type Farm = (typeof Farm)[keyof typeof Farm];

const RubberTreeVariety = {
    Burgundy: "Burgundy",
    Ruby: "Ruby",
    Tineke: "Tineke",
} as const;

type RubberTreeVariety = (typeof RubberTreeVariety)[keyof typeof RubberTreeVariety];

let myTree = new RubberTree("myTree", {type: RubberTreeVariety.Burgundy, farm: Farm.Codeinfra_Planters_Inc_})

export const myTreeType = myTree.type

export const myTreeFarmChanged = myTree.farm.apply(f => f + "foo");

export const mySentence = codeinfra.all([myTree.type, myTree.farm])
    .apply(([type, farm])=> `My ${type} Rubber tree is from ${farm}`)
