PROJECT_NAME     := Codeinfra Python SDK
LANGHOST_PKG     := github.com/khulnasoft/codeinfra/sdk/python/cmd/codeinfra-language-python/v3
VERSION          := $(if ${CODEINFRA_VERSION},${CODEINFRA_VERSION},$(shell ../../scripts/codeinfra-version.sh))
PYPI_VERSION 	   := $(if ${PYPI_VERSION},${PYPI_VERSION},$(shell ../../scripts/codeinfra-version.sh python))

ifeq ($(DEBUG),"true")
$(info    VERSION       = $(VERSION))
$(info    PYPI_VERSION  = $(PYPI_VERSION))
endif

PROJECT_PKGS    := $(shell go list -C cmd/codeinfra-language-python ./...)

include ../../build/common.mk

# Motivation: running `make TEST_ALL_DEPS= test_all` permits running
# `test_all` without the dependencies.
TEST_ALL_DEPS ?= build

ensure::
	@command -v uv > /dev/null || { echo "uv command not found. Please install it first." && exit 1; }
	uv venv --allow-existing
	uv sync --dev

build_package:: ensure
	uv run -m build --outdir ./build --installer uv

build_plugin::
	go install -C cmd/codeinfra-language-python \
		-ldflags "-X github.com/khulnasoft/codeinfra/sdk/v3/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}

build:: build_package build_plugin

lint:: ensure
	uv run -m ruff format --check
	MYPYPATH=./stubs uv run -m mypy ./lib/codeinfra
	uv run -m ruff check
	uv run -m pyright

lint_fix:: ensure
	make format
	make lint

format:: ensure
	uv run -m ruff format

install_package:: build_package
	cp ./cmd/codeinfra-language-python-exec "$(CODEINFRA_BIN)"
	cp ./dist/codeinfra-resource-codeinfra-python "$(CODEINFRA_BIN)"
	cp ./dist/codeinfra-analyzer-policy-python "$(CODEINFRA_BIN)"

install_plugin:: build_plugin
	GOBIN=$(CODEINFRA_BIN) go install -C cmd/codeinfra-language-python \
		  -ldflags "-X github.com/khulnasoft/codeinfra/sdk/v3/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}

install:: install_package install_plugin

test_go:: $(TEST_ALL_DEPS)
	@cd cmd/codeinfra-language-python && \
		$(GO_TEST) ${PROJECT_PKGS}

test_fast:: $(TEST_ALL_DEPS)
	. .venv/*/activate && ./scripts/test_fast.sh

test_auto:: $(TEST_ALL_DEPS)
	. .venv/*/activate && ./scripts/test_auto.sh

test_all:: test_fast test_auto test_go

dist::
	go install -C cmd/codeinfra-language-python \
		-ldflags "-X github.com/khulnasoft/codeinfra/sdk/v3/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}
	cp ./cmd/codeinfra-language-python-exec "$$(go env GOPATH)"/bin/
	cp ./dist/codeinfra-resource-codeinfra-python "$$(go env GOPATH)"/bin/
	cp ./dist/codeinfra-analyzer-policy-python "$$(go env GOPATH)"/bin/

brew:: BREW_VERSION := $(shell ../../scripts/get-version HEAD)
brew::
	go install -C cmd/codeinfra-language-python \
		-ldflags "-X github.com/khulnasoft/codeinfra/sdk/v3/go/common/version.Version=${BREW_VERSION}" ${LANGHOST_PKG}
	cp ./cmd/codeinfra-language-python-exec "$$(go env GOPATH)"/bin/
	cp ./dist/codeinfra-resource-codeinfra-python "$$(go env GOPATH)"/bin/
	cp ./dist/codeinfra-analyzer-policy-python "$$(go env GOPATH)"/bin/

publish::
	for file in ../../artifacts/sdk-python-*.whl; do \
		basename=$$(basename $${file}); \
		mv -vT "$${file}" "../../artifacts/$${basename##sdk-python-}"; \
	done
	twine upload \
	-u "${PYPI_USERNAME}" -p "${PYPI_PASSWORD}" \
		../../artifacts/*.whl \
		--skip-existing \
		--verbose
