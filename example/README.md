# DEMO

```console
$ cd example

$ cat Dockerfile
FROM busybox:1.34
RUN echo "Hello World!" >/banner

$ tug -t mcandre/tug-demo -load linux/amd64

$ docker run --rm mcandre/tug-demo cat /banner
Hello World!
```
