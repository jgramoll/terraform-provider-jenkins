package main

import (
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type ensureDefinitionFunc func(client.JobDefinition) error
type definitionCodeFunc func(client.JobDefinition) string
type definitionImportScriptFunc func(string, client.JobDefinition) string

var ensureDefinitionFuncs = map[string]ensureDefinitionFunc{}
var definitionCodeFuncs = map[string]definitionCodeFunc{}
var definitionImportScriptFuncs = map[string]definitionImportScriptFunc{}

func jobDefinitionCode(definition client.JobDefinition) string {
	if definition != nil {
		reflectType := reflect.TypeOf(definition).String()
		if codeFunc, ok := definitionCodeFuncs[reflectType]; ok {
			return codeFunc(definition)
		} else {
			log.Println("[WARNING] Unkown Job Definition:", reflectType)
		}
	}
	return ""
}

func jobDefinitionsImportScript(jobName string, definition client.JobDefinition) string {
	if definition != nil {
		reflectType := reflect.TypeOf(definition).String()
		if codeFunc, ok := definitionImportScriptFuncs[reflectType]; ok {
			return codeFunc(jobName, definition)
		} else {
			log.Println("[WARNING] Unkown Job Definition:", reflectType)
		}
	}
	return ""
}
