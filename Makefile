VERSION := $(shell cat VERSION)
LDFLAGS := -X 'cli-pass/cmd.Version=$(VERSION)'

build:
	go build -ldflags "$(LDFLAGS)" -o cli-pass .

release:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o cli-pass-linux
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o cli-pass-mac
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o cli-pass.exe