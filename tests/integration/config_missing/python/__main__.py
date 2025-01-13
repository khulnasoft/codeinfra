# Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import codeinfra

config = codeinfra.Config('config_missing_py')
config.require_secret('notFound')