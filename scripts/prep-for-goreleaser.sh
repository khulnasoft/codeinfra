#!/usr/bin/env bash

set -euo pipefail

# Populates ./bin directories for use by goreleaser.
rm -rf ./bin

LOCAL="${1:-"false"}"

COMMIT_TIME=$(git log -n1 --pretty='format:%cd' --date=format:'%Y%m%d%H%M')

install_file () {
    src="$1"
    shift

    for OS in "$@"; do # for each argument after the first:
        # if LOCAL == "true" and `go env goos` is not equal to the OS, skip it
        if [ "${LOCAL}" = "local" ] && [ "$(go env GOOS)" != "${OS}" ]; then
            continue
        fi
        DESTDIR="bin/${OS}"
        mkdir -p "${DESTDIR}"
        dest=$(basename "${src}")
        cp "$src" "${DESTDIR}/${dest}"
        touch -t "${COMMIT_TIME}" "$dest"
    done
}

install_file sdk/nodejs/dist/codeinfra-analyzer-policy                         linux   darwin
install_file sdk/nodejs/dist/codeinfra-analyzer-policy.cmd                     windows

install_file sdk/nodejs/dist/codeinfra-resource-codeinfra-nodejs                  linux   darwin
install_file sdk/nodejs/dist/codeinfra-resource-codeinfra-nodejs.cmd              windows

install_file sdk/python/dist/codeinfra-analyzer-policy-python                  linux   darwin
install_file sdk/python/dist/codeinfra-analyzer-policy-python.cmd              windows

install_file sdk/python/dist/codeinfra-resource-codeinfra-python                  linux   darwin
install_file sdk/python/dist/codeinfra-resource-codeinfra-python.cmd              windows

install_file sdk/python/dist/codeinfra-python-shim.cmd                         windows
install_file sdk/python/dist/codeinfra-python3-shim.cmd                        windows

install_file sdk/python/cmd/codeinfra-language-python-exec          linux darwin windows

# Get codeinfra-watch binaries
./scripts/get-codeinfra-watch.sh "${LOCAL}"
./scripts/get-language-providers.sh "${LOCAL}"
