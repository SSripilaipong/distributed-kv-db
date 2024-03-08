package main

import "embed"

//go:embed with_deps.go.tmpl
var resultTmplWithDeps embed.FS
