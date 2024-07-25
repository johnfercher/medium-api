GO_FILES = $(shell find . '(' -path '*/.*' -o -path './vendor' ')' -prune -o -name '*.go' -print | cut -b3-)
GO_PATHS =  $(shell go list -f '{{ .Dir }}' ./... | grep -E -v 'docs|cmd|mocks')

.PHONY: dod
dod: build test fmt lint

.PHONY: build
build:
	go build $(GO_PATHS)

.PHONY: test
test:
	go test $(GO_PATHS)

.PHONY: fmt
fmt:
	gofmt -s -w ${GO_FILES}
	gofumpt -l -w ${GO_FILES}
	goimports -w ${GO_PATHS}

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yml ./...
	make mock-lint

.PHONY: mock-lint
mock-lint:
	bash shell/mock-check.sh

.PHONY: install
install:
	bash shell/install.sh

.PHONY: docs
docs:
	docsify serve docs/

.PHONY: godoc
godoc:
	godoc -http=127.0.0.1:6060

.PHONY: mocks
mocks:
	rm -R mocks || true
	mockery
	make fmt

build-docker:
	docker build -t product-api .

run-docker:
	docker run -p 8081:8081 product-api
