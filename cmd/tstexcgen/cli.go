package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func readFlags() request {
	var nDeps int
	var nIns int
	var nOuts int
	var packageName string

	flag.IntVar(&nDeps, "dep", 0, "")
	flag.IntVar(&nIns, "in", 0, "")
	flag.IntVar(&nOuts, "out", 0, "")
	flag.StringVar(&packageName, "pkg", "", "")
	flag.Parse()
	path := flag.Args()[0]

	fileName := fmt.Sprintf("func_%ddep_%din_%dout.go", nDeps, nIns, nOuts)

	return request{
		PackageName: packageName,
		OutputPath:  filepath.Join(path, fileName),
		NDeps:       nDeps,
		NIns:        nIns,
		NOuts:       nOuts,
	}
}
