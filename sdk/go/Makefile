PROJECT_NAME     := Codeinfra Go SDK
LANGHOST_PKG     := github.com/khulnasoft/codeinfra/sdk/go/codeinfra-language-go/v3
VERSION          := $(if ${CODEINFRA_VERSION},${CODEINFRA_VERSION},$(shell ../../scripts/codeinfra-version.sh go))
TEST_FAST_PKGS   := $(shell go list ./codeinfra/... ./common/... | grep -v /vendor/ | grep -v templates)
TEST_AUTO_PKGS   := $(shell go list ./auto/... | grep -v /vendor/ | grep -v templates)

ifeq ($(DEBUG),"true")
$(info    VERSION         = $(VERSION))
endif

include ../../build/common.mk

# Motivation: running `make TEST_ALL_DEPS= test_all` permits running
# `test_all` without the dependencies.
TEST_ALL_DEPS ?= install

gen::
	go generate ./codeinfra/...

ensure: go.ensure

build:: gen
	go install -C codeinfra-language-go \
		-ldflags "-X github.com/khulnasoft/codeinfra/sdk/v3/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}

install_plugin::
	GOBIN=$(CODEINFRA_BIN) go install -C codeinfra-language-go \
		-ldflags "-X github.com/khulnasoft/codeinfra/sdk/v3/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}

install:: install_plugin

test_all:: test_fast test_auto

test_fast:: $(TEST_ALL_DEPS)
	@$(GO_TEST_FAST) $(TEST_FAST_PKGS)

	@cd codeinfra-language-go && \
		$(GO_TEST_FAST) $(shell cd codeinfra-language-go && go list ./... | grep -v /vendor/ | grep -v templates)

test_auto:: $(TEST_ALL_DEPS)
	@$(GO_TEST) $(TEST_AUTO_PKGS)

dist::
	go install -C codeinfra-language-go \
		-ldflags "-X github.com/khulnasoft/codeinfra/sdk/v3/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}

brew:: BREW_VERSION := $(shell ../../scripts/get-version HEAD)
brew::
	go install -C codeinfra-language-go \
		-ldflags "-X github.com/khulnasoft/codeinfra/sdk/v3/go/common/version.Version=${BREW_VERSION}" ${LANGHOST_PKG}

lint:: golangci-lint.ensure
	cd .. && golangci-lint run -c ../.golangci.yml --timeout 5m --path-prefix ..

	cd codeinfra-language-go && \
		golangci-lint run -c ../../../.golangci.yml --timeout 5m --path-prefix codeinfra-language-go

publish:
	git tag sdk/v${VERSION}
	git push origin sdk/v${VERSION}
