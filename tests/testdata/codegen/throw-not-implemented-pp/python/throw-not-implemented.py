import codeinfra


def not_implemented(msg):
    raise NotImplementedError(msg)

codeinfra.export("result", not_implemented("expression here is not implemented yet"))
