PROJECT := github.com/surajnarwade/kstatus
GITCOMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
PKGS := $(shell go list  ./... | grep -v $(PROJECT)/vendor)
BUILD_FLAGS := -ldflags="-w -X $(PROJECT)/cmd.GITCOMMIT=$(GITCOMMIT)"

default: bin

.PHONY: bin
bin:
	go build ${BUILD_FLAGS} -o status main.go
	cp status ~/.kube/plugins/status