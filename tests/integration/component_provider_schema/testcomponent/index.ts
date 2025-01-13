// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

class Provider implements codeinfra.provider.Provider {
    public readonly version = "0.0.1";
    constructor(public readonly schema?: string) {
    }
}

export function main(args: string[]) {
    const schema = process.env.INCLUDE_SCHEMA ? `{"hello": "world"}` : undefined;
    return codeinfra.provider.main(new Provider(schema), args);
}

main(process.argv.slice(2));
