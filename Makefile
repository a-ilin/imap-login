VERSION := $(shell git describe --tags --always)
ifndef VERSION
$(error VERSION not set (git issue?))
endif

LD_FLAGS := "-s -w -extldflags "-static" -X 'main.AppVersion=$(VERSION)'"

.PHONY: \
	imap-login \
	all \
	clean \
	install

all: \
	imap-login

imap-login:
	@ mkdir -p bin
	env CGO_ENABLED=0 go build -buildmode=pie -ldflags=$(LD_FLAGS) -o bin/imap-login .

clean:
	@ rm bin/imap-login > /dev/null 2>&1 || true

install: imap-login
	install -m 0755 -D -t /usr/local/bin bin/imap-login
	install -m 0644 -D LICENSE /usr/local/share/doc/imap-login/copyright
