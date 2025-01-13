#!/usr/bin/env bash

CODEINFRA_TEST_COVERAGE_PATH=$CODEINFRA_TEST_COVERAGE_PATH

set -euo pipefail

coverage run -m pytest lib/test/automation

if [[ "$CODEINFRA_TEST_COVERAGE_PATH" ]]; then
    if [ -e .coverage ]; then
        UUID=$(python -c "import uuid; print(str(uuid.uuid4()).replace('-', '').lower())")
        coverage xml -o $CODEINFRA_TEST_COVERAGE_PATH/python-auto-$UUID.xml
    fi
fi
