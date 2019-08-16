package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type ensureDefinitionFunc func(client.JobDefinition) error
type definitionCodeFunc func(client.JobDefinition) string
type definitionImportScriptFunc func(string, client.JobDefinition) string

var ensureDefinitionFuncs = map[string]ensureDefinitionFunc{}
var definitionCodeFuncs = map[string]definitionCodeFunc{}
var definitionImportScriptFuncs = map[string]definitionImportScriptFunc{}

func ensureJobDefinition(definition client.JobDefinition) error {
	if definition == nil {
		return nil
	}
	if definition.GetId() == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		definition.SetId(id.String())
	}
	reflectType := reflect.TypeOf(definition).String()
	ensureFunc, ok := ensureDefinitionFuncs[reflectType]
	if !ok {
		return fmt.Errorf("Unknown Definition Type %s", reflectType)
	}
	return ensureFunc(definition)
}

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
