package main

import (
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobBuildDiscarderPropertyStrategyCodeFunc func(string, client.JobBuildDiscarderPropertyStrategy) string
type jobBuildDiscarderPropertyStrategyImportScriptFunc func(string, string, client.JobBuildDiscarderPropertyStrategy) string

var jobBuildDiscarderPropertyStrategyCodeFuncs = map[string]jobBuildDiscarderPropertyStrategyCodeFunc{}
var jobBuildDiscarderPropertyStrategyImportScriptFuncs = map[string]jobBuildDiscarderPropertyStrategyImportScriptFunc{}

func jobBuildDiscarderPropertyStrategyCode(propertyIndex string, strategy client.JobBuildDiscarderPropertyStrategy) string {
	reflectType := reflect.TypeOf(strategy).String()
	if codeFunc, ok := jobBuildDiscarderPropertyStrategyCodeFuncs[reflectType]; ok {
		return codeFunc(propertyIndex, strategy)
	} else {
		log.Println("[WARNING] Unkown Job Build Discarder Property Strategy:", reflectType)
	}
	return ""
}

func jobBuildDiscarderPropertyStrategyImportScript(jobName string, propertyId string, strategy client.JobBuildDiscarderPropertyStrategy) string {
	reflectType := reflect.TypeOf(strategy).String()
	if codeFunc, ok := jobBuildDiscarderPropertyStrategyImportScriptFuncs[reflectType]; ok {
		return codeFunc(jobName, propertyId, strategy)
	} else {
		log.Println("[WARNING] Unkown Job Build Discarder Property Strategy:", reflectType)
	}
	return ""
}
