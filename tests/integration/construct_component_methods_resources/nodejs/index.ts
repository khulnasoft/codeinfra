// Copyright 2016-2021, KhulnaSoft Ltd.  All rights reserved.

import { Component } from "./component";

const component = new Component("component");

export const result = component.createRandom({ length: 10 }).result;
