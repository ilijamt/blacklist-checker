BUILD_DIR=build
BIN_DIR=bin
BIN_FILE=blacklist-checker
COVERAGE_DIR=${BUILD_DIR}/coverage
CONTRIB_DIR="contrib"
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
	${BIN_DIR}/${BIN_FILE} --completion-script-bash > ${CONTRIB_DIR}/.${BIN_FILE}.bash
	${BIN_DIR}/${BIN_FILE} --completion-script-zsh > ${CONTRIB_DIR}/.${BIN_FILE}.zsh
	${BIN_DIR}/${BIN_FILE} --help-man > ${CONTRIB_DIR}/${BIN_FILE}.1

.PHONY: install
install: build
	cp ${BIN_DIR}/${BIN_FILE} ${GOPATH}/bin/${BIN_FILE}
	cp ${CONTRIB_DIR}/.${BIN_FILE}.bash ${CONTRIB_DIR}/.${BIN_FILE}.zsh ${HOME}
	touch "${HOME}/.bash_completion"
	grep -q -F '[ -s "${HOME}/.${BIN_FILE}.bash" ] && . ${HOME}/.${BIN_FILE}.bash' ${HOME}/.bash_completion || echo '[ -s "${HOME}/.${BIN_FILE}.bash" ] && . ${HOME}/.${BIN_FILE}.bash' >> "${HOME}/.bash_completion"

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

.PHONY: package
package:
	debuild --preserve-env --preserve-envvar PATH -us -uc -d