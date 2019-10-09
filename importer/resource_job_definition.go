package main

import (
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type definitionCodeFunc func(client.JobDefinition) string

var definitionCodeFuncs = map[string]definitionCodeFunc{}

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
