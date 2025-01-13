# Copyright 2024, KhulnaSoft Ltd.

from codeinfra import ComponentResource, ResourceOptions

from echo import Echo

for i in range(0, 100):
    Echo(f"echo-{i}", echo=f"hello-{i}")
