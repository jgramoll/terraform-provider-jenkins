package main

import (
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobActionCodeFunc func(client.JobAction) string
type jobActionImportScriptFunc func(string, client.JobAction) string

var jobActionCodeFuncs = map[string]jobActionCodeFunc{}
var jobActionImportScriptFuncs = map[string]jobActionImportScriptFunc{}

func jobActionsCode(actions *client.JobActions) string {
	code := ""
	for _, item := range *actions.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobActionCodeFuncs[reflectType]; ok {
			code += codeFunc(item)
		} else {
			log.Println("[WARNING] Unknown action type:", reflectType)
		}
	}
	return code
}

func jobActionsImportScript(jobName string, actions *client.JobActions) string {
	code := ""
	for _, item := range *actions.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobActionImportScriptFuncs[reflectType]; ok {
			code += codeFunc(jobName, item)
		} else {
			log.Println("[WARNING] Unknown action type:", reflectType)
		}
	}
	return code
}
