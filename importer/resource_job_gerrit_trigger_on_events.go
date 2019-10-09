package main

import (
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerOnEventsCodeFunc func(string, client.JobGerritTriggerOnEvent) string
type jobGerritTriggerOnEventsImportScriptFunc func(
	string, string, string, client.JobGerritTriggerOnEvent) string

var jobGerritTriggerOnEventsCodeFuncs = map[string]jobGerritTriggerOnEventsCodeFunc{}
var jobGerritTriggerOnEventsImportScriptFuncs = map[string]jobGerritTriggerOnEventsImportScriptFunc{}

func jobGerritTriggerOnEventsCode(
	triggerIndex string, events *client.JobGerritTriggerOnEvents,
) string {
	code := ""
	for _, item := range *events.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobGerritTriggerOnEventsCodeFuncs[reflectType]; ok {
			code += codeFunc(triggerIndex, item)
		} else {
			log.Println("[WARNING] Unknown gerrit trigger on event type:", reflectType)
		}
	}
	return code
}

func jobGerritTriggerOnEventsImportScript(
	jobName string, propertyId string, triggerId string, events *client.JobGerritTriggerOnEvents,
) string {
	code := ""
	for _, item := range *events.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobGerritTriggerOnEventsImportScriptFuncs[reflectType]; ok {
			code += codeFunc(jobName, propertyId, triggerId, item)
		} else {
			log.Println("[WARNING] Unknown gerrit trigger on event type:", reflectType)
		}
	}
	return code
}
