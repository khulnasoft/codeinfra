// Copyright 2016-2021, KhulnaSoft Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import "mocha";
import * as assert from "assert";

import * as codeinfra from "@codeinfra/codeinfra";

import { listStorageAccountKeysOutput, ListStorageAccountKeysResult } from "../listStorageAccountKeys";
import * as amiIds from "../getAmiIds";
import { GetAmiIdsFilterArgs } from "../types/input";

codeinfra.runtime.setMocks({
    newResource: function(_: codeinfra.runtime.MockResourceArgs): {id: string, state: any} {
        throw new Error("newResource not implemented");
    },
    call: function(args: codeinfra.runtime.MockCallArgs) {
        if (args.token == "mypkg::listStorageAccountKeys") {
            return {
                "keys": [{
                    "creationTime": "my-creation-time",
                    "keyName": "my-key-name",
                    "permissions": "my-permissions",
                    "value": JSON.stringify(args.inputs),
                }]
            };
        }

        if (args.token == "mypkg::getAmiIds") {
            return {
                sortAscending: args.inputs.sortAscending,
                nameRegex: args.inputs.nameRegex,
                owners: args.inputs.owners,
                id: JSON.stringify({filters: args.inputs.filters})
            }
        }

        throw new Error("call not implemented for " + args.token);
    },
});

function checkTable(done: any, transform: (res: any) => any, table: {given: codeinfra.Output<any>, expect: any}[]) {
    checkOutput(done, codeinfra.all(table.map(x => x.given)), res => {
        res.map((actual, i) => {
            assert.deepStrictEqual(transform(actual), table[i].expect);
        });
    });
}

describe("output-funcs", () => {

    it("getAmiIdsOutput", (done) => {

        function filter(n: number): GetAmiIdsFilterArgs {
            return {
                name: codeinfra.output(`filter-${n}-name`),
                values: [
                    codeinfra.output(`filter-${n}-value-1`),
                    codeinfra.output(`filter-${n}-value-2`)
                ],
            }
        }

        const output = amiIds.getAmiIdsOutput({
            owners: [codeinfra.output("owner-1"),
                     codeinfra.output("owner-2")],
            nameRegex: codeinfra.output("[a-z]"),
            sortAscending: codeinfra.output(true),
            filters: [filter(1), filter(2)],
        });

        checkOutput(done, output, (res: amiIds.GetAmiIdsResult) => {
            assert.equal(res.sortAscending, true);
            assert.equal(res.nameRegex, "[a-z]");
            assert.deepStrictEqual(res.owners, ["owner-1", "owner-2"]);

            assert.deepStrictEqual(JSON.parse(res.id), {
                filters: [
                    {
                        name: 'filter-1-name',
                        values: [
                            'filter-1-value-1',
                            'filter-1-value-2'
                        ]
                    },
                    {
                        name: 'filter-2-name',
                        values: [
                            'filter-2-value-1',
                            'filter-2-value-2'
                        ]
                    },
                ]
            });
        });
    });

    it("listStorageAccountKeysOutput", (done) => {
        const output = listStorageAccountKeysOutput({
            accountName: codeinfra.output("my-account-name"),
            resourceGroupName: codeinfra.output("my-resource-group-name"),
        });
        checkOutput(done, output, (res: ListStorageAccountKeysResult) => {
            assert.equal(res.keys.length, 1);
            const k = res.keys[0];
            assert.equal(k.creationTime, "my-creation-time");
            assert.equal(k.keyName, "my-key-name");
            assert.equal(k.permissions, "my-permissions");
            assert.deepStrictEqual(JSON.parse(k.value), {
                "accountName": "my-account-name",
                "resourceGroupName": "my-resource-group-name"
            });
        });
    });

    it("listStorageAccountKeysOutput with optional arg set", (done) => {
        const output = listStorageAccountKeysOutput({
            accountName: codeinfra.output("my-account-name"),
            resourceGroupName: codeinfra.output("my-resource-group-name"),
            expand: codeinfra.output("my-expand"),
        });
        checkOutput(done, output, (res: ListStorageAccountKeysResult) => {
            assert.equal(res.keys.length, 1);
            const k = res.keys[0];
            assert.equal(k.creationTime, "my-creation-time");
            assert.equal(k.keyName, "my-key-name");
            assert.equal(k.permissions, "my-permissions");
            assert.deepStrictEqual(JSON.parse(k.value), {
                "accountName": "my-account-name",
                "resourceGroupName": "my-resource-group-name",
                "expand": "my-expand"
            });
        });
    });

 });


function checkOutput<T>(done: any, output: codeinfra.Output<T>, check: (value: T) => void) {
    output.apply(value => {
        try {
            check(value);
            done();
        } catch (error) {
            done(error);
        }
    });
}
