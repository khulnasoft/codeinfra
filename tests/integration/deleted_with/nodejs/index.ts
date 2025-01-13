// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import { FailsOnDelete } from "./failsOnDelete"
import { Random } from "./random"

const rand = new Random("random", { length: 10 });
new FailsOnDelete("failsondelete", { deletedWith: rand });
