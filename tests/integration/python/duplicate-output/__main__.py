# Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import codeinfra

a = codeinfra.Output.from_input([1, 2])

codeinfra.export("export1", a)
codeinfra.export("export2", a)