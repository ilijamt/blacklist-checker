BUILD_DIR=build
BIN_DIR=bin
BIN_FILE=blacklist-checker
COVERAGE_DIR=${BUILD_DIR}/coverage
COVERAGE_MODE=count
GOPACKAGES=./...
GOFILES_GLIDE=$(shell glide novendor)
GOFILES_NOVENDOR=$(shell find . -type f -name '*.go' -not -path "*/vendor/*")
VERSION_FILE=VERSION

BUILD_VERSION:=$(shell git log --pretty=format:'%h' -n 1)
BUILD_DATE:=$(shell date -u)
ifneq ($(wildcard $(VERSION_FILE)),)
	VERSION:=$(shell cat $(VERSION_FILE))
else
	VERSION:=
endif

.PHONY: build
build: 
	go build -ldflags "-X 'main.BuildVersion=${VERSION}' -X 'main.BuildHash=${BUILD_VERSION}' -X 'main.BuildDate=${BUILD_DATE}'" -o "${BIN_DIR}/${BIN_FILE}" .

.PHONY: install
install: 
	go build -ldflags "-X 'main.BuildVersion=${VERSION}' -X 'main.BuildHash=${BUILD_VERSION}' -X 'main.BuildDate=${BUILD_DATE}'" -o "${GOPATH}/bin/${BIN_FILE}" .

.PHONY: autocomplete
autocomplete: build
	bin/${BIN_FILE} --completion-script-bash > ${BIN_FILE}.bash
	bin/${BIN_FILE} --completion-script-zsh > ${BIN_FILE}.zsh

# Checks project and source code if everything is according to standard
.PHONY: check
check:
	gofmt -l ${GOFILES_NOVENDOR} | (! grep . -q) || (echo "Code differs from gofmt's style" && false)
	go vet ${GOFILES_GLIDE}

# Runs gofmt -w on the project's source code, modifying any files that do not
# match its style.
.PHONY: fmt
fmt:
	gofmt -l -w ${GOFILES_NOVENDOR}

# Runs gofmt -s -w on the project's source code, modifying any files that do not
# match its style.
.PHONY: simplify
simplify:
	gofmt -l -s -w ${GOFILES_NOVENDOR}

# Run the tests for the application
.PHONY: test
test:
	go test $(shell glide novendor) -v

.PHONY: lint
lint:
	@golint
