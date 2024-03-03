# Project
SHELL := /bin/bash
NAME := terraform-provider-vmware-fusion
PROJECT := terraform-provider-vmware-fusion
VERSION := 0.1.0
BUILD_DATE := $(shell date +%Y%m%d)
GIT_SHA := $(shell git rev-parse --short=8 HEAD)

# Toolchain
GO := GO111MODULE=on GOPROXY="https://goproxy.cn,direct" go
GO_VERSION := $(shell $(GO) version | sed -e 's/^[^0-9.]*\([0-9.]*\).*/\1/')

# Main
BINARY := terraform-provider-vmware-fusion_v$(VERSION)
MAIN := ./main.go
PACKAGE := github.com/leryn1122/terraform-provider-vmware-fusion

##@ General

.PHONY: help
help: ## Print help info.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Developement

.PHONY: install
install: ## Install dependencies.
	$(GO) get -d -v ./...

.PHONY: check
check: ## Check
	$(GO) vet ./...

.PHONY: fmt
fmt: ## Format against code.
	$(GO) fmt ./...

.PHONY: clean
clean: ## Clean target artifact.
	$(GO) clean -r -x

.PHONY: unittest
unittest: ## Run all unit tests.
	TF_ACC=1 $(GO) test ./... -v $(TESTARGS) -timeout 120m

.PHONY: test
test: ## Run all integrate tests.
	cd sample && terraform plan && terraform apply

##@ Build

.PHONY: build
build: ## Build target artifact.
	$(GO) build -a -ldflags '-extldflags "-static" ' \
	  -o target/$(BINARY) $(MAIN)
