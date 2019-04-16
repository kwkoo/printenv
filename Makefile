PREFIX=github.com/kwkoo
PACKAGE=printenv

GOPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GOBIN=$(GOPATH)/bin
IMAGENAME="kwkoo/$(PACKAGE)"
VERSION="0.1"

.PHONY: run build clean runcontainer newapp

run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOPATH)/src/$(PREFIX)/$(PACKAGE)/main.go

build:
	@echo "Building..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) CGO_ENABLED=0 GOOS=linux \
	  go build \
	  -a \
	  -installsuffix cgo \
	  -o $(GOBIN)/$(PACKAGE) \
	  $(PREFIX)/$(PACKAGE)

clean:
	rm -f $(GOPATH)/bin/$(PACKAGE) $(GOPATH)/pkg/*/$(PACKAGE).a

image: $(GOBIN)/$(PACKAGE)
	docker build --rm -t $(IMAGENAME):$(VERSION) $(GOPATH)

runcontainer:
	docker run --rm -it --name $(PACKAGE) $(IMAGENAME):$(VERSION)

newapp:
	oc new-app $(IMAGENAME):$(VERSION)
