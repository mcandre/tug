# tug: Multi-platform Docker rescue ship

![logo](tug.png)

# ABOUT

tug streamlines Docker pipelines.

Spend less time managing buildx images. Enjoy more time developing your core application.

# EXAMPLE

```console
$ cd example

$ tug -t mcandre/tug-demo -exclude-arch mips64

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

Why Docker?

Docker containerization technology has revolutionized software. Developers can enjoy easier development environment setup. Enormous amounts of disk space can be freed up when switching between different software projects. The entire build environment can be simulated accurately, ensuring new contributors can quickly iterate on the product. For users, cloud services running in Kubernetes containers exhibit a much higher degree of precision and reliability, compared with older, manual ways of system administration. Docker did not invent build reproducibility or container technology, but Docker's ease of use and the popular Docker Hub registry have paved the way for widespread adoption of modern software organization principles.

Why multi-platform?

Multi-platform images help developers to serve users across more user environments. For example, a developer who builds Docker images from a classic workstation (amd64) can more confidently support users consuming the images on a modern M1 Mac or Raspberry Pi (arm64). Multi-platform images are the glue that binds these software components together in a working fashion.

Why tug?

In particular, the docker buildx subsystem provides fundamental operations in support of managing multi-platform images. However, Docker has grown to a large, complicated supersystem requiring a significant amount of handholding for common tasks. This presents the opportunity for tug to come by and pick up the slack. tug simply chains together buildx primitives into larger, more practical workflows.

# DOCUMENTATION

https://godoc.org/github.com/mcandre/tug

# DOWNLOAD

https://github.com/mcandre/tug/releases

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/tug/cmd/tug@latest
```

# RUNTIME REQUIREMENTS

* [Docker](https://www.docker.com/) v20.10.12+

# CONTRIBUTING

For more information on developing tug itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

# LICENSE

FreeBSD

# USAGE

`tug -get-platforms` lists available platforms. Generally of the form `linux/*`.

`tug -ls <name>` lists any buildx cache entries present for the given image name, of the form `name[:tag]`.

`tug -t <name>` builds multi-platform images into the buildx cache, of the form `name[:tag]`.

* `-debug` enables additional logging. In case of some buildx error.
* `-exclude-os <list>` / `-exclude-arch <list>` skip the specified operating systems and/or architectures. For example, any platform variants unsupported by your `FROM` base image.
* `-load <platform>` copies an image to the local Docker registry as a side effect of the build. By default, Docker does not copy any buildx images to the local Docker registry as witnessed by `docker image`, `docker run`, etc. Select an appropriate `linux/<architecture>` platform based on your host machine. Typically `-load linux/amd64` for traditional hosts, or `-load linux/arm64` for newer arm64 hosts.
* `-push` uploads buildx cached images to the remote Docker registry, as a side effect of the image build process. This works around gaps in the buildx subsystem for conventional build, push workflows.
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
