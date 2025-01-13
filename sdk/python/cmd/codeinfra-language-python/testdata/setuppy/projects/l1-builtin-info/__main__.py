import codeinfra

codeinfra.export("stackOutput", codeinfra.get_stack())
codeinfra.export("projectOutput", codeinfra.get_project())
codeinfra.export("organizationOutput", codeinfra.get_organization())
