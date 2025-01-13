// Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

import { Component } from "./component"

const component = new Component("testComponent", {
    foo: "bar"
});

const message = component.getMessage({ echo: "hello" }).message;
