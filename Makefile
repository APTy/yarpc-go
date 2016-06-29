# Minor versions of Go for which the lint check should be run.
LINTABLE_MINOR_VERSIONS := 6

# Paths besides auto-detected generated files that should be excluded from
# lint results.
LINT_EXCLUDES_EXTRAS =

##############################################################################
export GO15VENDOREXPERIMENT=1

GO_VERSION := $(shell go version | cut -d' ' -f3)   # e.g.: go1.6.2
GO_MINOR_VERSION := $(word 2, $(subst ., , $(GO_VERSION)))

ifneq ($(filter $(LINTABLE_MINOR_VERSIONS),$(GO_MINOR_VERSION)),)
SHOULD_LINT := true
endif

PACKAGES := $(shell glide novendor)

GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

# Files whose first line contains "Code generated by" are generated.
GENERATED_GO_FILES := $(shell \
	find $(GO_FILES) \
	-exec sh -c 'head -n30 {} | grep "Code generated by\|\(Autogenerated\|Automatically generated\) by\|@generated" >/dev/null' \; \
	-print)

LINT_EXCLUDES := $(GENERATED_GO_FILES) $(LINT_EXCLUDES_EXTRAS)

# Pipe lint output into this to filter out ignored files.
FILTER_LINT := grep -v $(patsubst %,-e %, $(LINT_EXCLUDES))

##############################################################################

.PHONY: build
build:
	go build $(PACKAGES)

.PHONY: lint
lint:
ifdef SHOULD_LINT
	$(eval FMT_LOG := $(shell mktemp -t gofmt.XXXXX))
	@gofmt -e -s -l $(GO_FILES) | $(FILTER_LINT) > $(FMT_LOG) || true
	@[ ! -s "$(FMT_LOG)" ] || (echo "gofmt failed:" | cat - $(FMT_LOG) && false)

	$(eval VET_LOG := $(shell mktemp -t govet.XXXXX))
	@go vet $(PACKAGES) 2>&1 | grep -v '^exit status' | $(FILTER_LINT) > $(VET_LOG) || true
	@[ ! -s "$(VET_LOG)" ] || (echo "govet failed:" | cat - $(VET_LOG) && false)

	$(eval LINT_LOG := $(shell mktemp -t golint.XXXXX))
	@cat /dev/null > $(LINT_LOG)
	@$(foreach pkg, $(PACKAGES), golint $(pkg) | $(FILTER_LINT) >> $(LINT_LOG) || true;)
	@[ ! -s "$(LINT_LOG)" ] || (echo "golint failed:" | cat - $(LINT_LOG) && false)
else
	@echo "Skipping linters for $(GO_VERSION)"
endif

.PHONY: install
install:
	glide --version || go get github.com/Masterminds/glide
	glide install


.PHONY: test
test:
	go test $(PACKAGES)


.PHONY: cover
cover:
	./scripts/cover.sh $(shell go list $(PACKAGES))
	go tool cover -html=cover.out -o cover.html


# This is not part of the regular test target because we don't want to slow it
# down.
.PHONY: test-examples
test-examples:
	make -C examples


.PHONY: crossdock
crossdock:
	docker-compose kill go
	docker-compose rm -f go
	docker-compose build go
	docker-compose run crossdock


.PHONY: crossdock-fresh
crossdock-fresh: install
	docker-compose kill
	docker-compose rm --force
	docker-compose pull
	docker-compose build
	docker-compose run crossdock


.PHONY: install_ci
install_ci: install
ifdef SHOULD_LINT
	go get -u -f github.com/golang/lint/golint
endif
	go get github.com/wadey/gocovmerge
	go get github.com/mattn/goveralls
	go get golang.org/x/tools/cmd/cover

.PHONY: lint_ci
lint_ci: lint

.PHONY: test_ci
test_ci:
	./scripts/cover.sh $(shell go list $(PACKAGES))
