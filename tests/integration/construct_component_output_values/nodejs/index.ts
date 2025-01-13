// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";

import { Component } from "./component";

new Component("component", {
    foo: {
        something: "hello",
    },
    bar: {
        tags: {
            "a": "world",
            "b": codeinfra.secret("shh"),
        },
    },
});
