package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"log"
	"reflect"
)

type jobBuildDiscarderPropertyStrategyCodeFunc func(client.JobBuildDiscarderPropertyStrategy) string

var jobBuildDiscarderPropertyStrategyCodeFuncs = map[string]jobBuildDiscarderPropertyStrategyCodeFunc{}

func init() {
	jobPropertyCodeFuncs["*client.JobBuildDiscarderProperty"] = jobBuildDiscarderPropertyCode
}

func jobBuildDiscarderPropertyStrategyCode(strategy client.JobBuildDiscarderPropertyStrategy) string {
	reflectType := reflect.TypeOf(strategy).String()
	if codeFunc, ok := jobBuildDiscarderPropertyStrategyCodeFuncs[reflectType]; ok {
		return codeFunc(strategy)
	} else {
		log.Println("[WARNING] Unkown Job Build Discarder Property Strategy:", reflectType)
	}
	return ""
}
