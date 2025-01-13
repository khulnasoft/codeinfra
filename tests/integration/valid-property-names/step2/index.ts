// Copyright 2016-2023, KhulnaSoft Ltd.  All rights reserved.
import { Resource } from "./resource";

export const a = new Resource("a", {
    state: {
        template: {
            metadata: {
                annotations: {},
            },
        },
    }
});
