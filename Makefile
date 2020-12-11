#####################################################
# GNU Makefile
# @Author xczh <xczh.me@foxmail.com>
#####################################################
# Meta Info
BUILDVERSION ?= $(shell git describe --tags 2>/dev/null)
BUILDGITHASH := $(shell git rev-parse HEAD 2>/dev/null)
BUILDTIME := $(shell date "+%Y/%m/%d %H:%M:%S %Z(%z)")
BUILDGOVERSION := $(shell go version)
MAKE_VERSION := $(shell $(MAKE) -v | head -n 1)

# Config

# Enable UPX Compress
ENABLE_UPX ?=

# Set NO_DEBUG=1 will add '-s -w' to LDFLAGS
NO_DEBUG ?=

# Enable race detector. e.g ENABLE_RACE=1
ENABLE_RACE ?=

# Set fake GOROOT for trace log path. e.g. GOROOT_FINAL=/tmp/go
GOROOT_FINAL ?= /tmp/go

# Go Toolchain
export GO111MODULE=on
GO ?= go
CGO_ENABLED ?= 0
GO_OUTPUT := build/
#GOFILES := $(shell find . -name "*.go" -type f ! -path "*/bindata.go")
GOFMT ?= gofmt -s

# Generate LDFLAGS / GCFLAGS / ASMFLAGS
BUILD_OPTS := -trimpath -tags=jsoniter
LDFLAGS := -X 'main.BuildVersion=$(BUILDVERSION)' -X main.BuildGitHash=$(BUILDGITHASH) -X 'main.BuildTime=$(BUILDTIME)' -X 'main.BuildGoVersion=$(BUILDGOVERSION)' -X 'main.MakeVersion=$(MAKE_VERSION)'
GCFLAGS :=
ASMFLAGS :=

ifdef NO_DEBUG
    LDFLAGS += -s -w
endif

ifdef ENABLE_RACE
    BUILD_OPTS += -race
endif

# Cross Build Function
# $(1): Output Filename
# $(2): GOOS
# $(3): GOARCH
# $(4): go package or file
define go-build
	GOROOT_FINAL=$(GOROOT_FINAL) \
	CGO_ENABLED=$(CGO_ENABLED) \
	GOOS=$(2) \
	GOARCH=$(3) \
	$(GO) build \
	-o "$(GO_OUTPUT)$(1)" \
	-ldflags "$(LDFLAGS)" \
	-gcflags "$(GCFLAGS)" \
	-asmflags "$(ASMFLAGS)" \
	$(BUILD_OPTS) \
	$(4)
	@if [ -n "$(ENABLE_UPX)" ]; then \
	    upx $(GO_OUTPUT)$(1); \
	fi;
endef

#####################################################
# Target
#####################################################
all: clean build

clean:
	rm -rf build/app_*

build:
	$(call go-build,app_$(BUILDVERSION),,,./cmd/main.go)

# test:
# 	$(GO) test -v -cover ./cmd
run:
	$(GO) run ./cmd/main.go

convey:
	-goconvey -host 0.0.0.0 -port 8080 -launchBrowser false -packages 4 -excludedDirs "build,public,views,vendor,node_modules"

cross_build: cross_build_windows cross_build_linux cross_build_darwin

cross_build_windows: cross_build_windows_386 cross_build_windows_amd64

cross_build_windows_386:
	$(call go-build,app_$(BUILDVERSION)_windows_386.exe,windows,386,./cmd/main.go)

cross_build_windows_amd64:
	$(call go-build,app_$(BUILDVERSION)_windows_amd64.exe,windows,amd64,./cmd/main.go)

cross_build_linux: cross_build_linux_arm cross_build_linux_386 cross_build_linux_amd64

cross_build_linux_arm:
	$(call go-build,app_$(BUILDVERSION)_linux_arm,linux,arm,./cmd/main.go)

cross_build_linux_386:
	$(call go-build,app_$(BUILDVERSION)_linux_386,linux,386,./cmd/main.go)

cross_build_linux_amd64:
	$(call go-build,app_$(BUILDVERSION)_linux_amd64,linux,amd64,./cmd/main.go)

cross_build_darwin: cross_build_darwin_amd64

cross_build_darwin_amd64:
	$(call go-build,app_$(BUILDVERSION)_darwin_amd64,darwin,amd64,./cmd/main.go)

.PHONY: all clean build install test cross_build cross_build_windows cross_build_linux cross_build_darwin