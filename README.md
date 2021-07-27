# gocmt - Add missing comment on exported function, method, type, constant, variable in go file

![Build status](https://github.com/cuonglm/gocmt/actions/workflows/ci.yaml/badge.svg?branch=master)
[![Go Reference](https://pkg.go.dev/badge/github.com/cuonglm/gocmt.svg)](https://pkg.go.dev/github.com/cuonglm/gocmt)
[![Go Report Card](https://goreportcard.com/badge/github.com/cuonglm/gocmt)](https://goreportcard.com/report/github.com/cuonglm/gocmt)

# Installation
```sh
go get -u github.com/cuonglm/gocmt
```

# Why gocmt

Some of my projects have many files with exported fields, variables, functions missing comment, so lint tools will complain.

I find a way to auto add missing comment to them, just to pass the lint tools but nothing existed.

So `gocmt` comes in.

# Usage
```sh
$ gocmt -h
usage: gocmt [flags] [file ...]
  -d string
    	Directory to process
  -i	Make in-place editing
  -t string
    	Comment template (default "...")
```

# Example
```sh
$ cat testdata/main.go
package p

var i = 0

var I = 1

var c = "constant un-exported"

const C = "constant exported"

type t struct{}

type T struct{}

func main() {
}

func unexport(s string) {
}
func Export(s string) {
}

func ExportWithComment(s string) {
}
```

Using `gocmt` give you:
```sh
$ gocmt testdata/main.go
package p

var i = 0

// I ...
var I = 1

var c = "constant un-exported"

// C ...
const C = "constant exported"

type t struct{}

// T ...
type T struct{}

func main() {
}

func unexport(s string) {
}
// Export ...
func Export(s string) {
}

// ExportWithComment ...
func ExportWithComment(s string) {
}
```

Default template is `...`, you can change it using `-t` option.

# Author

Cuong Manh Le <cuong.manhle.vn@gmail.com>

# License

See [LICENSE](https://github.com/cuonglm/gocmt/blob/master/LICENSE)
