package main

import (
	"fmt"
	"os"

	"github.com/dstull/csv2json/model"
)

func main() {
	enc, err := model.FromArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Initial error: %v\n", err)
		os.Exit(3)
	}
	if err = enc.Encode(); err != nil {
		fmt.Fprintf(os.Stderr, "Run time error: %v\n", err)
		os.Exit(1)
	}
}
