# Copyright 2016-2024, KhulnaSoft Ltd.  All rights reserved.

"""An example program that type checks with pyright but pyright is not installed"""

import codeinfra

# This export won't work because the first argument is a number, not a string
codeinfra.export(42, 'bar')
