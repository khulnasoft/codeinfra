// Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import * as codeinfra from "@codeinfra/codeinfra";
import { Resource } from "./resource";

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);

// "a" should be in the checkpoint.
