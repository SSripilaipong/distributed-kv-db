package main

import (
	"bytes"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
	"fmt"
	"go/format"
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
	if len(detail.Deps) > 0 {
		return renderWithDeps(detail)
	}
	return renderWithoutDeps(detail)
}

var renderWithDeps = fn.Compose(
	fn.Compose(rslt.WrapErrorFunc[string]("renderTemplateToString fail"), renderTemplateToString(resultTmplWithDeps)),
	funcDetailToRenderDetail,
)

var renderWithoutDeps = fn.Compose(
	fn.Compose(rslt.WrapErrorFunc[string]("renderTemplateToString fail"), renderTemplateToString(resultTmplWithDeps)),
	funcDetailToRenderDetail,
)

func funcDetailToRenderDetail(detail funcDetail) renderDetail {
	allInputVars := append(detail.Deps, detail.Ins...)
	allVars := append(allInputVars, detail.Outs...)

	return renderDetail{
		funcDetail:             detail,
		RenderDefaultVars:      allInputVars,
		AllGenericTypeNamesStr: strings.Join(genericTypeNames(allVars), ","),
		DepGenericTypeNamesStr: strings.Join(genericTypeNames(detail.Deps), ","),
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
		var buff bytes.Buffer
		err = t.Execute(&buff, detail)
		if err != nil {
			return rslt.Error[string](fmt.Errorf("render error: %w", err))
		}
		raw := buff.String()
		resultBytes, err := format.Source([]byte(raw))
		if err != nil {
			return rslt.Error[string](fmt.Errorf("format error: %w", err))
		}
		return rslt.Value(string(resultBytes))
	}
}
