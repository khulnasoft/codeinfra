// Copyright 2016-2022, KhulnaSoft Ltd.

import * as codeinfra from '@codeinfra/codeinfra'

// A ResourceProvider using the default generic type, with explicit return type defined.
class DefaultGenericProvider implements codeinfra.dynamic.ResourceProvider {
  async create (props: any): Promise<codeinfra.dynamic.CreateResult> {
    return { id: 'resource-id', outs: {} }
  }

  async check (olds: any, news: any): Promise<codeinfra.dynamic.CheckResult> {
    return Promise.resolve({ inputs: news })
  }

  async diff (id: codeinfra.ID, olds: any, news: any): Promise<codeinfra.dynamic.DiffResult> {
    return Promise.resolve({})
  }

  async delete (id: codeinfra.ID, props: any): Promise<void> { return Promise.resolve() }

  async update (id: codeinfra.ID, olds: any, news: any): Promise<codeinfra.dynamic.UpdateResult> {
    return Promise.resolve({ outs: {} })
  }

  async read(id: codeinfra.ID, props: any): Promise<codeinfra.dynamic.ReadResult> {
    return Promise.resolve({ props: {} })
  }
}

type InputArgs = {
  names: string
}

type OutputArgs = {
  resourceId: string
  name: string
}

// All parameters and returns typed are inferred through the generic types provided.
class TypedGenericProvider implements codeinfra.dynamic.ResourceProvider<InputArgs, OutputArgs> {
  async create (props) {
    return { id: 'resource-id', outs: { resourceId: "id", name: "test" } }
  }

  async check (olds, news) {
    return Promise.resolve({ inputs: news })
  }

  async diff (id, olds, news) {
    return Promise.resolve({})
  }

  async delete (id, props) { return Promise.resolve() }

  async update (id, olds, news) {
    return Promise.resolve({ outs: { resourceId: olds.resourceId, ...news } })
  }

  async read(id, props) {
    return Promise.resolve({ props: props })
  }
}
