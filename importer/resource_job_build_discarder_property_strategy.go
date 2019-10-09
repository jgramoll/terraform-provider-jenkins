package main

import (
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobBuildDiscarderPropertyStrategyCodeFunc func(client.JobBuildDiscarderPropertyStrategy) string

var jobBuildDiscarderPropertyStrategyCodeFuncs = map[string]jobBuildDiscarderPropertyStrategyCodeFunc{}

func jobBuildDiscarderPropertyStrategyCode(strategy client.JobBuildDiscarderPropertyStrategy) string {
	reflectType := reflect.TypeOf(strategy).String()
	if codeFunc, ok := jobBuildDiscarderPropertyStrategyCodeFuncs[reflectType]; ok {
		return codeFunc(strategy)
	} else {
		log.Println("[WARNING] Unkown Job Build Discarder Property Strategy:", reflectType)
	}
	return ""
}
