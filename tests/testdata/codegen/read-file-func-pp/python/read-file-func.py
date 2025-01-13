import codeinfra

key = (lambda path: open(path).read())("key.pub")
codeinfra.export("result", key)
