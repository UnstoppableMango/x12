GO     ?= go
GINKGO ?= $(GO) tool ginkgo

test:
	$(GINKGO) -r .

tidy:
	$(GO) mod tidy
