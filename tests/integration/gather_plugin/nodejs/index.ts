import * as codeinfra from "@codeinfra/codeinfra";

class Random extends codeinfra.Resource {
  result!: codeinfra.Output<string | undefined>;

  constructor(name: string, length: number, opts?: codeinfra.ResourceOptions) {
    const inputs: any = {};
    inputs["length"] = length;
    super("testprovider:index:Random", name, true, inputs, opts);
  }
}

class RandomProvider extends codeinfra.ProviderResource {
  constructor(name: string, opts?: codeinfra.ResourceOptions) {
    super("testprovider", name, {}, opts);
  }
}

const r = new Random("default", 10, {
  pluginDownloadURL: "get.example.test",
});
export const defaultProvider = r.result;

const provider = new RandomProvider("explicit", {
  pluginDownloadURL: "get.codeinfra.test/providers",
});

new Random("explicit", 8, { provider: provider });
