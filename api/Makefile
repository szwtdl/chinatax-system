SHELL=/usr/bin/env bash
NAME := chinatax
VERSION ?= prod
all: amd64 win mac

amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -ldflags "-X main.Version=$(VERSION)" -o ../build/cmd/$(NAME)-linux-$(VERSION) main.go
.PHONY: amd64

arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GOARM=7 go build -ldflags "-X main.Version=$(VERSION)" -o ../build/cmd/$(NAME)-linux-$(VERSION) main.go
.PHONY: arm64

win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 /usr/local/go/bin/go build -ldflags "-X main.Version=$(VERSION)" -o ../build/cmd/$(NAME).exe main.go
.PHONY: win

mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ../build/cmd/$(NAME)-darwin-$(VERSION) main.go
.PHONY: mac

clean:
	find ../build/cmd -maxdepth 1 -type f -name "$(NAME)-*" -delete
.PHONY: clean

