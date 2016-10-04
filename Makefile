GO       = go
GOX      = gox
GOX_ARGS = "-osarch=linux/amd64 linux/386 linux/arm linux/arm64 darwin/amd64 freebsd/amd64 freebsd/386 windows/386 windows/amd64 -output=build/{{.Dir}}_{{.OS}}_{{.Arch}}"

APP = nats_exporter
DIR = $(shell pwd)

NO_COLOR    = \033[0m
OK_COLOR    = \033[32;01m
ERROR_COLOR = \033[31;01m
WARN_COLOR  = \033[33;01m
MAKE_COLOR  = \033[33;01m%-20s\033[0m

SRCS      = $(shell git ls-files '*.go' | grep -v '^vendor/')
BUILD_DIR = build/

.DEFAULT_GOAL := build

.PHONY: help
help:
	@echo "$(OK_COLOR)==== $(APP) ====$(NO_COLOR)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(MAKE_COLOR) : %s\n", $$1, $$2}'

.PHONY: clean
clean:
	@echo "$(OK_COLOR)[$(APP)] Cleanup$(NO_COLOR)"
	@rm $(BUILD_DIR)/*

.PHONY: init
init:
	@echo "$(OK_COLOR)[$(APP)] Install requirements$(NO_COLOR)"
	@go get -u github.com/golang/glog
	@go get -u github.com/Masterminds/rmvcsdir
	@go get -u github.com/golang/lint/golint
	@go get -u github.com/kisielk/errcheck
	@go get -u golang.org/x/tools/cmd/oracle
	@go get -u github.com/mitchellh/gox
	@go get -u github.com/kardianos/govendor

.PHONY: deps-list
deps-list:
	@echo "$(OK_COLOR)[$(APP)] List dependencies$(NO_COLOR)"
	@govendor list

.PHONY: deps-update
deps-update:
	@echo "$(OK_COLOR)[$(APP)] Update dependencies$(NO_COLOR)"
	@govendor update

.PHONY: build
build:
	@echo "$(OK_COLOR)[$(APP)] Build $(NO_COLOR)"
	@$(GO) build -o $(BUILD_DIR)/nats_exporter .

.PHONY: test
test:
	@echo "$(OK_COLOR)[$(APP)] Launch unit tests $(NO_COLOR)"
	@govendor test +local

.PHONY: lint
lint:
	@$(foreach file,$(SRCS),golint $(file) || exit;)

.PHONY: vet
vet:
	@$(foreach file,$(SRCS),$(GO) vet $(file) || exit;)

.PHONY: release-build
release-build:
	@echo "$(OK_COLOR)[$(APP)] Create binaries $(NO_COLOR)"
	@go get -u github.com/mitchellh/gox
	@$(GOX) $(GOX_ARGS) github.com/lovoo/nats_exporter
