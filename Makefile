# Extract the latest Git tag as the version
VERSION := $(shell git describe --tags --abbrev=0)
LDFLAGS := -X 'vault/cmd.Version=$(VERSION)'

build:
	go build -ldflags "$(LDFLAGS)" -o vault .

release:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o vault-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o vault-darwin-amd64
	GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o vault-linux-arm64
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o vault-darwin-arm64
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o vault-windows-amd64.exe