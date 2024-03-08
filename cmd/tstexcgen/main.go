package main

import (
	"distributed-kv-db/common/zd"
	"fmt"
)

type request struct {
	NDeps       int
	NIns        int
	NOuts       int
	OutputPath  string
	PackageName string
}

type funcDetail struct {
	PackageName string
	Name        string
	Deps        []genericVarDetail
	Ins         []genericVarDetail
	Outs        []genericVarDetail
}

type genericVarDetail struct {
	Name            string
	VarName         string
	GenericTypeName string
}

func main() {
	err := processRequest(readFlags())
	if err != nil {
		panic(err)
	}
}

func processRequest(r request) error {
	name := fmt.Sprintf("Func%dDep%dIn%dOut", r.NDeps, r.NIns, r.NOuts)
	if r.NDeps <= 0 {
		name = fmt.Sprintf("Func%dIn%dOut", r.NIns, r.NOuts)
	}

	detail := funcDetail{
		PackageName: r.PackageName,
		Name:        name,
		Deps:        genDeps(r.NDeps),
		Ins:         genIns(r.NIns),
		Outs:        genOuts(r.NOuts),
	}
	result := renderFuncDetail(detail)
	if result.IsError() {
		return result.Error()
	}
	return writeFile(r.OutputPath, result.Value())
}

func genOuts(n int) (result []genericVarDetail) {
	for i := range zd.RangeCh(1, n+1) {
		result = append(result, genericVarDetail{
			Name:            fmt.Sprintf("Out%d", i),
			VarName:         fmt.Sprintf("o%d", i),
			GenericTypeName: fmt.Sprintf("O%d", i),
		})
	}
	return
}

func genIns(n int) (result []genericVarDetail) {
	for i := range zd.RangeCh(1, n+1) {
		result = append(result, genericVarDetail{
			Name:            fmt.Sprintf("In%d", i),
			VarName:         fmt.Sprintf("i%d", i),
			GenericTypeName: fmt.Sprintf("I%d", i),
		})
	}
	return
}

func genDeps(n int) (result []genericVarDetail) {
	for i := range zd.RangeCh(1, n+1) {
		result = append(result, genericVarDetail{
			Name:            fmt.Sprintf("Dep%d", i),
			VarName:         fmt.Sprintf("D%d", i),
			GenericTypeName: fmt.Sprintf("D%d", i),
		})
	}
	return
}
