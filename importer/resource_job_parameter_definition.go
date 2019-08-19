package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type ensureJobParameterDefinitionFunc func(client.JobParameterDefinition) error
type jobParameterDefinitionCodeFunc func(string, string, client.JobParameterDefinition) string
type jobParameterDefinitionImportScriptFunc func(string, string, string, client.JobParameterDefinition) string

var ensureJobParameterDefinitionFuncs = map[string]ensureJobParameterDefinitionFunc{}
var jobParameterDefinitionCodeFuncs = map[string]jobParameterDefinitionCodeFunc{}
var jobParameterDefinitionImportScriptFuncs = map[string]jobParameterDefinitionImportScriptFunc{}

func jobParameterDefinitionsCode(propertyIndex string, definitions *client.JobParameterDefinitions) string {
	code := ""

	for i, item := range *definitions.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobParameterDefinitionCodeFuncs[reflectType]; ok {
			definitionIndex := fmt.Sprintf("%v_%v", propertyIndex, i+1)
			code += codeFunc(propertyIndex, definitionIndex, item)
		} else {
			log.Println("[WARNING] Unknown parameter definition type:", reflectType)
		}
	}

	return code
}

func jobParameterDefinitionsImportScript(propertyIndex string, jobName string, propertyId string, definitions *client.JobParameterDefinitions) string {
	code := ""

	for i, item := range *definitions.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobParameterDefinitionImportScriptFuncs[reflectType]; ok {
			definitionIndex := fmt.Sprintf("%v_%v", propertyIndex, i+1)
			code += codeFunc(definitionIndex, jobName, propertyId, item)
		} else {
			log.Println("[WARNING] Unknown parameter definition type:", reflectType)
		}
	}

	return code
}
