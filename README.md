# tug: Multi-platform Docker rescue ship

![logo](tug.png)

# ABOUT

tug streamlines Docker pipelines.

Spend less time managing buildx images. Enjoy more time developing your core application.

# EXAMPLE

```console
$ cd example

$ tug -t mcandre/tug-demo

$ tug -ls mcandre/tug-demo
Platform:  linux/386
Platform:  linux/amd64
Platform:  linux/amd64/v2
...

$ tug -t mcandre/tug-demo -load linux/amd64

$ docker run --rm mcandre/tug-demo cat /banner
Hello World!
```

# MOTIVATION

buildx is hard. tug is easy.

When Docker introduced the buildx subsystem, their goals included making buildx operationally successful. But not necessarily as straightforward, consistent, and intuitive as single-platform `docker` commands. (Assuming that you consider Docker *straightforward, consistent, and intuitive*, ha.) We have run extensive drills on what buildx has to offer, and wrapped this into a neat little package called tug.

We are not replacing buildx, we just provide a proven workflow for high level buildx operation. We hope tug helps you to jumpstart multi-platform projects and even learn some fundamental buildx commands along the way. If you're curious to see more buildx gears turning, apply the `tug -debug` flag.

# DOCUMENTATION

https://godoc.org/github.com/mcandre/tug

# DOWNLOAD

https://github.com/mcandre/tug/releases

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/tug/cmd/tug@latest
```

# RUNTIME REQUIREMENTS

* [Docker](https://www.docker.com/) 20.10.12+

# CONTRIBUTING

For more information on developing tug itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

# LICENSE

FreeBSD

# USAGE

`tug -get-platforms` lists available platforms. Generally of the form `linux/*`.

`tug -ls <name>` lists any buildx cache entries present for the given image name, of the form `name[:tag]`.

`tug -t <name>` builds multi-platform images into the buildx cache, of the form `name[:tag]`.

* `-debug` enables additional logging. In case of some buildx error.
* `-platforms <list>` enables additional platforms. By default, tug targets all available supported platforms, minus any exceedingly niche platforms; See `-get-platforms`. The list is space delimited.
* `-exclude-os <list>` / `-exclude-arch <list>` skip the specified operating systems and/or architectures. For example, any variants unsupported by your `FROM` base image. The list is space delimited.
* `-load <platform>` copies an image to the local Docker registry as a side effect of the build. By default, Docker does not copy any buildx images to the local Docker registry as witnessed by `docker image`, `docker run`, etc. Select an appropriate `linux/<architecture>` platform based on your host machine. Typically `-load linux/amd64` for traditional hosts, or `-load linux/arm64` for newer arm64 hosts.
* `-push` uploads buildx cached images to the remote Docker registry, as a side effect of the image build process. This works around gaps in the buildx subsystem for conventional build, push workflows.
* `-extra <list>` sends additional command line flags to `docker buildx build` commands. The list is comma delimited.
* `.` or `<directory>` are optional trailing arguments for the Docker build directory. We default to the current working directory.

`tug -clean` empties the buildx image cache and removes the `tug` builder.

See `tug -help` for more detail.

# tug-in-docker

Running tug itself within a Docker context, such as for CI/CD, would naturally require Docker-in-Docker privileges. See the relevant documentation for your particular cluster environment, such as Kubernetes.

# DOCKER HUB COMMUNITY

[Docker Hub](https://hub.docker.com/) provides an exceptional variety of base images, everything from Debian to Ubuntu to RHEL to glibc to musl to uClibC. If your base image lacks support for a particular platform, try searching for alternative base images. Or, build a new base image from scratch and publish it back to Docker Hub! The more we refine our base images, the easier it is to extend and use them.

# SEE ALSO

* [factorio](https://github.com/mcandre/factorio) ports Go applications.
* [gox](https://github.com/mitchellh/gox), an older Go cross-compiler wrapper.
* [LLVM](https://llvm.org/) bitcode offers an abstract assembler format for C/C++ code.
* [snek](https://github.com/mcandre/snek) ports native C/C++ applications.
* [tonixxx](https://github.com/mcandre/tonixxx) ports applications of any programming language.
* [WASM](https://webassembly.org/) provides a portable interface for C/C++ code.
* [xgo](https://github.com/karalabe/xgo) supports Go projects with native cgo dependencies.
