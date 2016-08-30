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

# Run the coverage tests for the application
.PHONY: test-coverage
test-coverage:
	rm -vrf ${COVERAGE_DIR} ${BUILD_DIR}/coverage.html
	mkdir ${COVERAGE_DIR}/profiles -p
	for pkg in `go list $(shell glide novendor)`; do \
		cvrfile=$$(basename "$$pkg"); \
		go test --covermode=${COVERAGE_MODE} --coverprofile="${COVERAGE_DIR}/profiles/$$cvrfile.profile" "$$pkg"; \
	done
	$(MAKE) test-coverage-assemble

.PHONY: test-coverage-assemble
test-coverage-assemble:
ifneq ($(wildcard $(COVERAGE_DIR)/profiles/*.profile),)
	touch ${COVERAGE_DIR}/coverprofile
	echo "mode: ${COVERAGE_MODE}" > ${COVERAGE_DIR}/coverprofile
	grep -h -v "^mode:" ${COVERAGE_DIR}/profiles/*.profile >> ${COVERAGE_DIR}/coverprofile
	go tool cover -html=${COVERAGE_DIR}/coverprofile -o "${BUILD_DIR}/coverage.html"
else
	rm -vrf ${BUILD_DIR}
	@echo "No profiles generated, no tests found"
endif

.PHONY: lint
lint:
	@golint
	@golint queue
