GO       ?= go
GINKGO   ?= $(GO) tool ginkgo
GOCILINT ?= $(GO) tool golangci-lint

test:
	$(GINKGO) -r .

lint:
	$(GOCILINT) run

tidy:
	$(GO) mod tidy
