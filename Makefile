VERSION := "$(shell git describe --abbrev=0)"
MODULE := "$(shell git config --get remote.origin.url | sed 's|^https\://\([^ <]*\)\(.*\)\.git|\1|g')"
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(dir $(MKFILE_PATH))
update-pkg-cache:
	cd $$HOME && \
	GOPROXY=https://proxy.golang.org GO111MODULE=on go get $(MODULE)@$(VERSION)

gen:
	go generate ./...

test:
	go test ./...	
	go test ./... -tags=safe