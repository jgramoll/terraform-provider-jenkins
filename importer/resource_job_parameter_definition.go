package main

import (
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobParameterDefinitionCodeFunc func(client.JobParameterDefinition) string

var jobParameterDefinitionCodeFuncs = map[string]jobParameterDefinitionCodeFunc{}

func jobParameterDefinitionsCode(definitions *client.JobParameterDefinitions) string {
	code := ""

	for _, item := range *definitions.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobParameterDefinitionCodeFuncs[reflectType]; ok {
			code += codeFunc(item)
		} else {
			log.Println("[WARNING] Unknown parameter definition type:", reflectType)
		}
	}

	return code
}
