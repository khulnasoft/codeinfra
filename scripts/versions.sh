#!/usr/bin/env bash

set -euo pipefail

SCRIPTDIR=$(dirname "$0")
CODEINFRA_VERSION=$("${SCRIPTDIR}/codeinfra-version.sh")
echo GENERIC_VERSION="${CODEINFRA_VERSION}"
echo VERSION="${CODEINFRA_VERSION}"
echo PYPI_VERSION="$("${SCRIPTDIR}/codeinfra-version.sh" python)"
echo DOTNET_VERSION="$("${SCRIPTDIR}/codeinfra-version.sh" dotnet)"
echo GORELEASER_CURRENT_TAG="v${CODEINFRA_VERSION}"
