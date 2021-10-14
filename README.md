# FlipIP

Just a little quick app to take an IP (or set of IPs), and convert them to
reverse IP pointers

Example:

```bash
$ flipip 1.2.3.4
4.3.2.1.in-addr.arpa.
```

Intended to be used with things like [dog](https://github.com/ogham/dog) which
provide that sweet sweet dns tooling, but are missing reverse lookup helpers
