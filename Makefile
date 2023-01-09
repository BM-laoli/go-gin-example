.PHONY: build  tool lint help

all: build
build:
# 直接交叉编译
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o go-gin-example .

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

# clean:
# 	rm -rf go-gin-example
# 	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"