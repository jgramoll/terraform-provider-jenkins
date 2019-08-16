package main

import (
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type ensureJobParameterDefinitionFunc func(client.JobParameterDefinition) error
type jobParameterDefinitionCodeFunc func(int, int, client.JobParameterDefinition) string
type jobParameterDefinitionImportScriptFunc func(int, int, string, string, client.JobParameterDefinition) string

var ensureJobParameterDefinitionFuncs = map[string]ensureJobParameterDefinitionFunc{}
var jobParameterDefinitionCodeFuncs = map[string]jobParameterDefinitionCodeFunc{}
var jobParameterDefinitionImportScriptFuncs = map[string]jobParameterDefinitionImportScriptFunc{}

func jobParameterDefinitionsCode(propertyIndex int, definitions *client.JobParameterDefinitions) string {
	code := ""

	for i, item := range *definitions.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobParameterDefinitionCodeFuncs[reflectType]; ok {
			code += codeFunc(propertyIndex, i+1, item)
		} else {
			log.Println("[WARNING] Unknown parameter definition type:", reflectType)
		}
	}

	return code
}

func jobParameterDefinitionsImportScript(propertyIndex int, jobName string, propertyId string, definitions *client.JobParameterDefinitions) string {
	code := ""

	for i, item := range *definitions.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobParameterDefinitionImportScriptFuncs[reflectType]; ok {
			code += codeFunc(propertyIndex, i+1, jobName, propertyId, item)
		} else {
			log.Println("[WARNING] Unknown parameter definition type:", reflectType)
		}
	}

	return code
}
