// Copyright 2016-2023, KhulnaSoft Ltd.

import * as codeinfra from '@codeinfra/codeinfra'

class CustomResource extends codeinfra.dynamic.Resource {
  constructor (name: string, opts?: codeinfra.ResourceOptions) {
    super(new DummyResourceProvider(), name, {}, opts, "custom-provider", "CustomResource")
  }
}

class DummyResourceProvider implements codeinfra.dynamic.ResourceProvider {
  async create (props: any): Promise<codeinfra.dynamic.CreateResult> {
    var config = JSON.parse(process.env.CODEINFRA_CONFIG)

    return { id: config["id"], outs: {} }
  }
}

const resource = new CustomResource('resource-name')

export const rid = resource.id
