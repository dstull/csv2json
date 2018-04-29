# csv2json

[![Go Report Card](https://goreportcard.com/badge/github.com/dstull/csv2json?style=flat-square)](https://goreportcard.com/report/github.com/dstull/csv2json)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/dstull/csv2json)
[![Release](https://img.shields.io/github/release/dstull/csv2json.svg?style=flat-square)](https://github.com/dstull/csv2json/releases/latest)
[![CircleCI](https://circleci.com/gh/dstull/csv2json/tree/master.svg?style=shield)](https://circleci.com/gh/dstull/csv2json/tree/master)

Simple tool for converting CSVs to JSON

## Installation

First install [Go](http://golang.org).

```shell
GOBIN=$(pwd) GOPATH=$(mktemp -d) go get github.com/dstull/csv2json
go build
```

## Usage

```shell
./csv2json -h
Usage of csv2json:
  -i string
    	Input file (default: stdin)
  -o string
    	Output file (default: stdout)

$ more test.csv
first,second,third,fourth
a,b,c,d
e,f,g,h

$ ./csv2json -i=test.csv| jsonpp
{
  "columns": [
    "first",
    "second",
    "third",
    "fourth"
  ],
  "lines": [
    [
      "a",
      "b",
      "c",
      "d"
    ],
    [
      "e",
      "f",
      "g",
      "h"
    ]
  ]
}
```
