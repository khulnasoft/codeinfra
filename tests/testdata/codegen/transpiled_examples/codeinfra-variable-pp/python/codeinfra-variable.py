import codeinfra
import os

codeinfra.export("cwd", os.getcwd())
codeinfra.export("stack", codeinfra.get_stack())
codeinfra.export("project", codeinfra.get_project())
codeinfra.export("organization", codeinfra.get_organization())
