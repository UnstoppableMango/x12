GO     ?= go
GINKGO ?= $(GO) tool ginkgo

tidy:
	$(GO) mod tidy

test:
	$(GINKGO) -r .
