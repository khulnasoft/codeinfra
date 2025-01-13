// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import { Config } from "@codeinfra/codeinfra";

const config = new Config();
export const testVar = config.require("test");
