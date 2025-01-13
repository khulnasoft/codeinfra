# Copyright 2024, KhulnaSoft Ltd.  All rights reserved.

import codeinfra
import codeinfra_random as random


def transform(args):
    return codeinfra.ResourceTransformResult(
        props=args.props,
        opts=args.opts,
    )


random.RandomInteger(
    "test",
    max=999,
    min=100,
    opts=codeinfra.ResourceOptions(
        transforms=[transform],
    ),
)
