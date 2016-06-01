# URLTracer
Traces and prints the redirect path of a URL

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
