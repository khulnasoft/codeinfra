// This tests the runtime's ability to be loaded side-by-side with another copy of the same runtime library.
// This is a hard and subtle problem because the runtime is configured with a bunch of state, like whether
// we are doing a dry-run and, more importantly, RPC addresses to communicate with the engine.  Normally we
// go through the startup shim to configure all of these things, but when the second copy gets loaded we don't.
// Subsequent copies of the runtime are able to configure themselves by using environment variables.

let assert = require("assert");
let path = require("path");

const sdkPath = "../../../../../";

// Load the first copy:
let codeinfra1 = require(sdkPath);

// Now delete the entries in the require cache, and load up the second copy:
const resolvedSdkPath = path.dirname(require.resolve(sdkPath));
Object.keys(require.cache).forEach((path) => {
    if (path.startsWith(resolvedSdkPath)) {
        delete require.cache[path];
    }
});
let codeinfra2 = require(sdkPath);

// Make sure they are different:
assert(codeinfra1 !== codeinfra2, "codeinfra1 !== codeinfra2");
assert(codeinfra1.runtime !== codeinfra2.runtime, "codeinfra1.runtime !== codeinfra2.runtime");

// Check that various settings are equal:
assert.strictEqual(
    codeinfra1.runtime.isDryRun(),
    codeinfra2.runtime.isDryRun(),
    "codeinfra1.runtime.isDryRun() !== codeinfra2.runtime.isDryRun()",
);
assert.strictEqual(
    codeinfra1.runtime.getProject(),
    codeinfra2.runtime.getProject(),
    "codeinfra1.runtime.getProject() !== codeinfra2.runtime.getProject()",
);
assert.strictEqual(
    codeinfra1.runtime.getStack(),
    codeinfra2.runtime.getStack(),
    "codeinfra1.runtime.getStack() !== codeinfra2.runtime.getStack()",
);
assert.deepStrictEqual(
    codeinfra1.runtime.allConfig(),
    codeinfra2.runtime.allConfig(),
    "codeinfra1.runtime.allConfig() !== codeinfra2.runtime.getStack()",
);

// Check that the two runtimes agree on the stack resource
let stack1 = codeinfra1.runtime.getStackResource();
let stack2 = codeinfra2.runtime.getStackResource();
assert.strictEqual(stack1, stack2, "codeinfra1.runtime.getStackResource() !== codeinfra2.runtime.getStackResource()");

// allConfig should have caught this, but let's check individual config values too.
let cfg1 = new codeinfra1.Config("sxs");
let cfg2 = new codeinfra2.Config("sxs");
assert.strictEqual(cfg1.get("message"), cfg2.get("message"));

// Try and set a stack transformation
function transform1(args) {
    args.props["runtime1"] = 1;
    return { props: args.props, opts: args.opts };
}
function transform2(args) {
    args.props["runtime2"] = 2;
    return { props: args.props, opts: args.opts };
}

codeinfra1.runtime.registerStackTransformation(transform1);
codeinfra2.runtime.registerStackTransformation(transform2);

// Now do some useful things that require RPC connections:
codeinfra1.log.info("logging via Codeinfra1 works!");
codeinfra2.log.info("logging via Codeinfra2 works too!");
let res1 = new codeinfra1.CustomResource("test:x:resource", "p1p1p1");
res1.urn.apply((urn) => assert.strictEqual(urn, "test:x:resource::p1p1p1"));
let res2 = new codeinfra2.CustomResource("test:y:resource", "p2p2p2");
res2.urn.apply((urn) => assert.strictEqual(urn, "test:y:resource::p2p2p2"));

// Both resources should have the stack transforms applied
res1.runtime1.apply((value) => assert.strictEqual(value, 1));
res1.runtime2.apply((value) => assert.strictEqual(value, 2));
res2.runtime1.apply((value) => assert.strictEqual(value, 1));
res2.runtime2.apply((value) => assert.strictEqual(value, 2));

codeinfra1.runtime.registerResourcePackage("test1", {
    version: "0.0.1",
});
codeinfra2.runtime.registerResourcePackage("test2", {
    version: "0.0.2",
});
let test1codeinfra1 = codeinfra1.runtime.getResourcePackage("test1");
assert.strictEqual(test1codeinfra1.version, "0.0.1");
let test1codeinfra2 = codeinfra2.runtime.getResourcePackage("test1");
assert.strictEqual(test1codeinfra2.version, "0.0.1");
let test2codeinfra1 = codeinfra1.runtime.getResourcePackage("test2");
assert.strictEqual(test2codeinfra1.version, "0.0.2");
let test2codeinfra2 = codeinfra2.runtime.getResourcePackage("test2");
assert.strictEqual(test2codeinfra2.version, "0.0.2");

// Check that we can set mocks successfully
// We don't need an actual test monitor here, just something to set and get.
class Mocks {
    newResource(args) {
        return {
            id: args.inputs.name + "_id",
            state: {
                ...args.inputs,
            },
        };
    }
    call(args) {
        return args;
    }
}

let mocks = new Mocks();

codeinfra1.runtime.setMocks(mocks);
assert.strictEqual(codeinfra1.runtime.getMonitor().mocks, mocks);
assert.strictEqual(codeinfra2.runtime.getMonitor().mocks, mocks);
assert.strictEqual(codeinfra1.runtime.getMonitor(), codeinfra2.runtime.getMonitor());
assert(codeinfra1.runtime.hasMonitor());
assert(codeinfra2.runtime.hasMonitor());

codeinfra1.runtime._reset();
codeinfra2.runtime._reset();
