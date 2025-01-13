# Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

"""An example program that type checks with mypy"""

import codeinfra

# This export won't work because the first argument is a number, not a string
codeinfra.export(42, 'bar')
