import * as codeinfra from "@codeinfra/codeinfra";
import * as dynamic from "@codeinfra/codeinfra/dynamic";

class ReflectProvider implements dynamic.ResourceProvider {
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: codeinfra.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: codeinfra.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: inputs }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}

export class ReflectResource<T> extends dynamic.Resource {
    public readonly value!: codeinfra.Output<T>;

    constructor(name: string, value: codeinfra.Input<T>, opts?: codeinfra.CustomResourceOptions) {
        super(new ReflectProvider(), name, {value: value}, opts);
    }
}

class DummyProvider implements dynamic.ResourceProvider {
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: codeinfra.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: codeinfra.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: {"value": "hello"} }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}

export class DummyResource extends dynamic.Resource {
    public readonly value!: codeinfra.Output<string>;

    constructor(name: string, opts?: codeinfra.CustomResourceOptions) {
        super(new DummyProvider(), name, {}, opts);
    }
}
