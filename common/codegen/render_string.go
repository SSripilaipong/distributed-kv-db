package codegen

import (
	"bytes"
	"distributed-kv-db/common/rslt"
	"fmt"
	"go/format"
	"text/template"
)

func RenderString(t *template.Template, obj any) rslt.Of[string] {
	var buff bytes.Buffer
	err := t.Execute(&buff, obj)
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
