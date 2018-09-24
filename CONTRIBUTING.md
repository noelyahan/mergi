# Contributing Guide

* [Style Guide](#style-guide)
* [Build](#build)
* [Features Requested](#features-requested)


## Style Guide
### Code format
Mergi uses [gofmt](https://golang.org/cmd/gofmt) to format the code, you must use [gofmt](https://golang.org/cmd/gofmt) to format your code before submitting.

### linter
Mergi recommend using [golint](https://github.com/golang/lint) or [gometalinter](https://github.com/alecthomas/gometalinter) to check your code format.


## Build

Make sure that this folder is in `GOPATH`, then:

```bash
$ go build ./cmd/mergi
```

### Mergi needs you ! :)