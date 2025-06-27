# DEMO

```console
$ cd example

$ cat Dockerfile
FROM tianon/toybox:0.8
RUN echo "Hello World!" >/banner

$ tug -t mcandre/tug-demo -load linux/amd64

$ docker run --rm mcandre/tug-demo cat /banner
Hello World!
```

When publishing images, apply any relevant comma separated (`,`) skip patterns with `-exclude-os` and/or `-exclude-arch`.
