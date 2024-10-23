# imap-login

`imap-login` attempts to login to IMAP server with the credentials provided via command line or environment variables.

The tool may be used to validate the credentials, or to trigger an external authorization sequence.

Upon successful login the program executes the logout command and exits.

## Usage

```sh
$ bin/imap-login -h
NAME:
   imap-login - Test login to IMAP server

USAGE:
   imap-login [global options]

VERSION:
   v0.1

GLOBAL OPTIONS:
   --verbose, --vv             Verbose logging (default: false)
   --server value, -s value    Server name or IP [$IMAP_SERVER]
   --port value, -p value      Port (default: 143/993) [$IMAP_PORT]
   --tls, -t                   Use TLS (default: false) [$IMAP_TLS]
   --username value, -u value  Username [$IMAP_USERNAME]
   --password value, -P value  Password [$IMAP_PASSWORD]
   --help, -h                  show help
   --version, -v               print the version
```

## Installation

Prerequisites:

- Golang 1.22
- GNU Make 4.3

Build the application with the command:

```sh
make
```

Install the application to `/usr/local/bin` with the command (optional):

```sh
make install
```

### Deploying to Alpine Linux

On Alpine Linux it may be necessary to install glibc compatibility libraries in order to run the program:

```sh
apk add --no-cache gcompat libc6-compat
```

## Legal

License terms are specified in the `LICENSE` file.
