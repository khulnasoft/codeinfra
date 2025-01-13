// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import { Component } from "./component"

const component = new Component("component");
const result = component.getMessage({ echo: "hello" });
export const message = result.message;
