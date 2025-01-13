# Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.

import codeinfra

from fails_on_delete import FailsOnDelete
from random_ import Random

rand = Random("random", length=10)
FailsOnDelete("failsondelete", opts=codeinfra.ResourceOptions(deleted_with=rand))
