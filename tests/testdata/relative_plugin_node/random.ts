// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

export class Random extends codeinfra.Resource {
    result!: codeinfra.Output<string | undefined>;
  
    constructor(name: string, length: number, opts?: codeinfra.ResourceOptions) {
      const inputs: any = {};
      inputs["length"] = length;
      super("testprovider:index:Random", name, true, inputs, opts);
    }
  }