package main

import (
	"distributed-kv-db/common/codegen"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
	"fmt"
	"io/fs"
	"strings"
	"text/template"
)

type renderDetail struct {
	funcDetail
	RenderDefaultVars      []genericVarDetail
	AllGenericTypeNamesStr string
	DepGenericTypeNamesStr string
	InGenericTypeNamesStr  string
	OutGenericTypeNamesStr string
}

func renderFuncDetail(detail funcDetail) rslt.Of[string] {
	return render(detail)
}

var render = fn.Compose(
	fn.Compose(rslt.WrapErrorFunc[string]("renderTemplateToString fail"), renderTemplateToString(resultTmpl)),
	funcDetailToRenderDetail,
)

func funcDetailToRenderDetail(detail funcDetail) renderDetail {
	allVars := append(detail.Ins, detail.Outs...)

	return renderDetail{
		funcDetail:             detail,
		RenderDefaultVars:      detail.Ins,
		AllGenericTypeNamesStr: strings.Join(genericTypeNames(allVars), ","),
		InGenericTypeNamesStr:  strings.Join(genericTypeNames(detail.Ins), ","),
		OutGenericTypeNamesStr: strings.Join(genericTypeNames(detail.Outs), ","),
	}
}

var genericTypeNames = slc.MapFunc(genericVarDetailToGenericTypeName)

func genericVarDetailToGenericTypeName(d genericVarDetail) string {
	return d.GenericTypeName
}

func renderTemplateToString(tmplFs fs.FS) func(detail renderDetail) rslt.Of[string] {
	t, err := template.ParseFS(tmplFs, "*")

	return func(detail renderDetail) rslt.Of[string] {
		if err != nil {
			return rslt.Error[string](fmt.Errorf("parse template error: %w", err))
		}
		return codegen.RenderString(t, detail)
	}
}
