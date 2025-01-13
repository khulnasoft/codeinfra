import codeinfra

config = codeinfra.Config()
cidr_block = config.get("cidrBlock")
if cidr_block is None:
    cidr_block = "Test config variable"
codeinfra.export("cidrBlock", cidr_block)
