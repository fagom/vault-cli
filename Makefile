VERSION := $(shell cat VERSION)
LDFLAGS := -X 'vault/cmd.Version=$(VERSION)'

build:
	go build -ldflags "$(LDFLAGS)" -o vault .

release:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o vault-linux
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o vault-mac
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o vault.exe