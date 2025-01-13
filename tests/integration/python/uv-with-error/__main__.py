# Copyright 2024, KhulnaSoft Ltd.  All rights reserved.

import codeinfra_random as random

r = random.RandomString("random", length=16)

# This will fail because the name is already used
r2 = random.RandomString("random", length=16)
