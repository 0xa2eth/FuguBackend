GOCMD		=		go
GORUN		=		$(GOCMD) run
GOBUILD		=		CGO_ENABLED=0 $(GOCMD) build
DIST		=		./build
BINARY		=		$(DIST)

# 64-Bit
# Linux
TARGET_LINUX=		GOOS=linux GOARCH=amd64
# MacOS
TARGET_MACOS=		GOOS=darwin GOARCH=amd64
# windows
TARGET_WINDOWS=		GOOS=windows GOARCH=amd64
# freebsd
TARGET_FREEBSD=		GOOS=freebsd GOARCH=amd64

#GIT_COMMIT=$(shell git rev-list -1 HEAD)
GO_VERSION	=		$(shell go version)
BUILD_DATE	=		$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

LDFLAGS		=	-ldflags "-w -s"
GCFLAGS		=   -gcflags "-N -l"
#LDFLAGS=-ldflags "-w -s \
#-X 'gitlab.deeply.cn/doc/gocommon/build.GitCommit=$(GIT_COMMIT)' \
#-X 'gitlab.deeply.cn/doc/gocommon/build.GoVersion=$(GO_VERSION)' \
#-X 'gitlab.deeply.cn/doc/gocommon/build.BuildDate=$(BUILD_DATE)'"
.PHONY: build-windows
build-windows:
	mkdir -p $(DIST);
	$(TARGET_WINDOWS) $(GOBUILD) -o $(DIST)/fugu-win ./

.PHONY: build-mac
build-mac:
	mkdir -p $(DIST);
	$(TARGET_MACOS) $(GOBUILD) -o $(DIST)/fugu-mac ./
.PHONY: build-linux
build-linux:
	mkdir -p $(DIST);
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY) ./cmd/main.go

.PHONY: build-linux-mini
build-linux-mini:
	mkdir -p $(DIST);
	$(TARGET_LINUX) $(GOBUILD) $(LDFLAGS) -o $(BINARY) ./

.PHONY: build-all
build-all:build-windows build-linux build-mac

.PHONY: run
run:
	$(GORUN) $(DIST)/fugu

.PHONY: clean
clean:
	@if [ -f ./$(BINARY) ] ; then rm ./${BINARY}; fi
	@if [ -f ./$(DIST)/fugu-win ] ; then rm ./$(DIST)/fugu-win; fi
	@if [ -f ./$(DIST)/fugu-mac ] ; then rm ./$(DIST)/fugu-mac; fi

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: help
help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build-linux - 编译 Go 代码, 生成二进制文件(linux)"
	@echo "make build-linux-mini - 编译 Go 代码, 生成缩小版的二进制文件(linux)"
	@echo "make build-all - 编译 Go 代码，生成多平台的二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make tidy - 执行go mod tidy"