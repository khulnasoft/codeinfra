#!/usr/bin/env bash
#
# To be used with Goreleaser as `gobinary` implementation as a
# replacement for the `go` toolchain.
#
# Function: coverage-enabled builds for Codeinfra CLI
#
# This builds binaries via `go test -c` workaround. Disabled for
# Windows builds. Only enabled on the Codeinfra CLI binaries.

set -euo pipefail

CODEINFRA_TEST_COVERAGE_PATH=${CODEINFRA_TEST_COVERAGE_PATH:-}
CODEINFRA_BUILD_MODE=${CODEINFRA_BUILD_MODE:-}

COVER_PACKAGES=( \
    "github.com/khulnasoft/codeinfra/pkg/v3/..." \
    "github.com/khulnasoft/codeinfra/sdk/v3/..." \
    "github.com/khulnasoft/codeinfra/sdk/go/codeinfra-language-go/v3/..." \
    "github.com/khulnasoft/codeinfra/sdk/nodejs/cmd/codeinfra-language-nodejs/v3/..." \
    "github.com/khulnasoft/codeinfra/sdk/python/cmd/codeinfra-language-python/v3/..." \
)

# Join COVER_PACKAGES with commas.
COVERPKG=$(IFS=,; echo "${COVER_PACKAGES[*]}")

case "$1" in
    build)
        MODE="$CODEINFRA_BUILD_MODE"
        if [ -z "$MODE" ]; then
            # If a build mode was not specified,
            # guess based on whether a coverage path was supplied.
            MODE=coverage
            if [ -z "$CODEINFRA_TEST_COVERAGE_PATH" ]; then
                MODE=normal
            fi
        fi

        RACE=
        CGO_ENABLED=0
        if [ "$CODEINFRA_ENABLE_RACE_DETECTION" = "true" ]; then
            RACE='-race'

            if [ "$(go env GOOS)" != "darwin" ]; then
                # On macOS, we don't need CGO but windows and linux still do.
                CGO_ENABLED=1
            fi
        fi
        export CGO_ENABLED

        case "$MODE" in
            normal)
                shift
                go build ${RACE} "$@"
                ;;
            coverage)
                shift
                go build ${RACE} -cover -coverpkg "$COVERPKG" "$@"
                ;;
            *)
                echo "unknown build mode: $MODE"
                exit 1
                ;;
        esac
        ;;
    install)
        echo "install command is not supported, please use build"
        exit 1
        ;;
    *)
        go "$@"
        ;;
esac
