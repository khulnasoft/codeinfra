#!/usr/bin/env bash

set -euo pipefail

codeinfra version

time codeinfra destroy --yes

codeinfra config set mode new
time codeinfra up --yes --skip-preview

codeinfra config set mode alias
time codeinfra up --yes --skip-preview


export PATH=~/.codeinfra-dev/bin:$PATH

codeinfra version

time codeinfra destroy --yes

codeinfra config set mode new
time codeinfra up --yes --skip-preview

codeinfra config set mode alias
time codeinfra up --yes --skip-preview
