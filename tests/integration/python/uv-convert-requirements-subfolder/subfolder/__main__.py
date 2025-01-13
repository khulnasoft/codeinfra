# Copyright 2024, KhulnaSoft Ltd.  All rights reserved.

"""A Python project that uses the uv toolchain where the requirements.txt is in a parent folder."""

import codeinfra

codeinfra.export("foo", "bar")
