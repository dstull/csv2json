# csv2json

[![Go Report Card](https://goreportcard.com/badge/github.com/dstull/csv2json?style=flat-square)](https://goreportcard.com/report/github.com/dstull/csv2json)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/dstull/csv2json)
[![Release](https://img.shields.io/github/release/dstull/csv2json.svg?style=flat-square)](https://github.com/dstull/csv2json/releases/latest)

Simple tool for converting CSVs to JSON

## Installation

First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```shell
GOBIN=$(pwd) GOPATH=$(mktemp -d) go get github.com/baltimore-sun-data/csv2json
```

## Usage

```shell
$ csv2json
Usage of csv2json:
  -dest string
        Destination file (default: stdout)
  -src string
        Source file (default: stdin)

$ more test.csv
a,b,c
1,2,3

$ csv2json | jsonpp
[
        {
                "a": "1",
                "b": "2",
                "c": "3"
        }
]
```
