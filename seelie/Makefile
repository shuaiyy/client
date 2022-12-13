PACKAGES = $(shell find . -type d -not -path './.git*' -not -path './vendor*')
RELEASE_VERSION := $(shell cat VERSION)
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
TIME := $(shell date +"%Y-%m-%d.%H:%M:%S")
DAY := $(shell date +"%Y-%m-%d")
GIT_COUNT 	= $(shell git rev-list --all --count)
VERSION     = $(RELEASE_VERSION)-$(GIT_COUNT)-$(COMMIT)

GOOS ?= $(shell uname -s | tr [:upper:] [:lower:])
GOARCH ?= $(shell go env GOARCH)

# todo server/cli 版本管理尚未在代码里实现
FLAGS = -ldflags "-X main.VERSION=$(VERSION) -X $(SETTINGS).BuildCommit=$(COMMIT) -X $(SETTINGS).BuildDate=$(TIME)"

init-git-hook:
	@bash ./script/git/init.sh

cli:
	@echo [Build] version=$(VERSION) commit=$(COMMIT) date=$(TIME)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(FLAGS) -o seelie ./cmd/cli/main.go

os_list := darwin linux
os_arch := amd64 arm64
cli-all:
	# https://go.dev/doc/install/source#environment
	@rm -rf ./tarball/seelie-client*
	@for goos in $(os_list); do \
  		for goarch in $(os_arch); do \
  		  			echo [Build Client] $$goos $$goarch version=$(VERSION) commit=$(COMMIT) date=$(TIME); \
  		  			mkdir -p ./tarball/seelie-client-$$goos-$$goarch; \
  		  			GOOS=$$goos GOARCH=$$goarch go build $(FLAGS) -o ./tarball/seelie-client-$$goos-$$goarch/seelie ./cmd/cli/main.go; \
  		  			cp -rf ./tarball/README.md ./tarball/install.sh ./tarball/seelie-client-$$goos-$$goarch/; \
  		  			echo "tar -zcf ./tarball/seelie-client-$$goos-$$goarch-$(DAY).tar.gz -C tarball/seelie-client-$$goos-$$goarch/ seelie README.md"; \
  		  			tar -zcf ./tarball/seelie-client-$$goos-$$goarch-$(DAY).tar.gz -C tarball/seelie-client-$$goos-$$goarch/ seelie README.md install.sh; \
  		done; \
  	done
	@mv ./tarball/seelie-client-darwin-amd64-$(DAY).tar.gz ./tarball/seelie-client-macos-intel-x86-amd64-darwin-$(DAY).tar.gz
	@mv ./tarball/seelie-client-darwin-arm64-$(DAY).tar.gz ./tarball/seelie-client-macos-silicon-m1m2-arm64-darwin-$(DAY).tar.gz

lint:
	# golangci-lint run -D structcheck -D unused -D errcheck --timeout=6m todo golangci-lint暂不支持golang 1.18泛型, 先禁用struct检查
	golangci-lint run --timeout=6m ./...