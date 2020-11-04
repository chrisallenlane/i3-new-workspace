makefile := $(realpath $(lastword $(MAKEFILE_LIST)))

# paths
GOBIN  :=

# executables
COLUMN := column
GO     := go
GREP   := grep
LINT   := revive
RM     := rm
SED    := sed
SORT   := sort
bin    := i3-new-workspace

## build: compile the executable
.PHONY: build
build:
	$(GO) build -o $(bin) main.go

## setup: install revive (linter)
.PHONY: setup
setup:
	$(GO) get -u github.com/mgechev/revive 

## fmt: run go fmt
.PHONY: fmt
fmt:
	$(GO) fmt ./...

## lint: lint go source files
.PHONY: lint
lint:
	$(LINT) ./...

## vet: vet go source files
.PHONY: vet
vet:
	$(GO) vet ./...

## check: format, lint, and vet
.PHONY: check
check: | fmt lint vet

## clean: remove compiled executables
.PHONY: clean
clean:
	$(RM) -f $(bin)

## help: display this help text
.PHONY: help
help:
	@$(CAT) $(makefile) | \
	$(SORT)             | \
	$(GREP) "^##"       | \
	$(SED) 's/## //g'   | \
	$(COLUMN) -t -s ':'
