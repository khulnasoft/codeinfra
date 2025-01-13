# Copyright 2024, KhulnaSoft Ltd.  All rights reserved.

"""A Python project that uses the uv toolchain, without specifing the location for the virtualenv."""

import codeinfra

codeinfra.export("foo", "bar")
