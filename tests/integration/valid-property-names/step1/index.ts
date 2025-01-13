// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.
import * as codeinfra from "@codeinfra/codeinfra";
import { Resource } from "./resource";


let config = new codeinfra.Config();
export const a = new Resource("a", {
    state: {
        // Driven by table tests in steps_test.go.
        [config.require("propertyName")]: "foo",
    }
});
