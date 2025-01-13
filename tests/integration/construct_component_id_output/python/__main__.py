# Copyright 2016-2018, KhulnaSoft Ltd.  All rights reserved.

import codeinfra
from component import Component

component_a = Component("a", id="hello")

codeinfra.export("id", component_a.id)