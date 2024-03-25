package main

import "embed"

//go:embed main.go.tmpl
var resultTmpl embed.FS
