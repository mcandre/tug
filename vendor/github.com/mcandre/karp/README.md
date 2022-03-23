# karp - universal app launcher

# EXAMPLES

```console
$ karp karp.gif
...

$ karp https://raw.githubusercontent.com/mcandre/karp/master/karp.txt
...

$ karp
Usage: karp <path>

-v Show version information
```

![magikarp](https://raw.githubusercontent.com/mcandre/karp/master/karp.gif)

# ABOUT

`karp` is a cross-operating system binary for easily launching file-based applications, using the skratchdot/open-golang library to route calls to `xdg-open`, etc.

`karp` will use the default application registered for a given path's protocol, mimetype, or file extension.

# DOCUMENTATION

https://godoc.org/github.com/mcandre/karp

# DOWNLOAD

https://github.com/mcandre/karp/releases

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/karp/...@latest
```

# LICENSE

FreeBSD

# RUNTIME REQUIREMENTS

(None)

# CONTRIBUTING

For more information on developing karp itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

# CREDIT

Powered by [open-golang](https://github.com/skratchdot/open-golang)
