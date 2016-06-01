# urltrace
Traces and prints the redirect path of a URL

## Installation

In the bin directory there are precompiled binaries for linux and Darwin (OS X) systems. If you prefer, you can also build it yourself with the following command:

```
go get -u github.com/kkirsche/urltrace
cd $GOPATH/src/github.com/kkirsche/urltrace
go install -race
```

## Usage
`urltrace` is designed to allow a user to trace the redirect path of a URL and record that so that they can identify any URLs which are necessary to reach a given URL. The command may be used like so:

```
Usage:
  urltrace [flags]

Flags:
  -f, --full-url      Display the entire URL, not the host portion.
  -t, --timeout int   Sets the timeout in seconds for a requested URL (default 10)
```

## Usage Examples
```
urltrace http://www.google.com/mail

urltrace --timeout 15 http://www.google.com/mail

urltrace -t 15 http://www.google.com/mail

urltrace --timeout 15 --full-url http://www.google.com/mail

urltrace -t 15 -f http://www.google.com/mail
```
