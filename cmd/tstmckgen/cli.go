package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func readFlags() request {
	var nIns int
	var nOuts int
	var packageName string

	flag.IntVar(&nIns, "in", 0, "")
	flag.IntVar(&nOuts, "out", 0, "")
	flag.StringVar(&packageName, "pkg", "", "")
	flag.Parse()
	path := flag.Args()[0]

	fileName := fmt.Sprintf("func_%din_%dout.go", nIns, nOuts)

	return request{
		PackageName: packageName,
		OutputPath:  filepath.Join(path, fileName),
		NIns:        nIns,
		NOuts:       nOuts,
	}
}
