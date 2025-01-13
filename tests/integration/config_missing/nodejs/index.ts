// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import { Config } from "@codeinfra/codeinfra";

const config = new Config("config_missing_js");
config.requireSecret("notFound")