package main

import (
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerOnEventsCodeFunc func(
	int, int, client.JobGerritTriggerOnEvent) string
type jobGerritTriggerOnEventsImportScriptFunc func(
	string, string, string, client.JobGerritTriggerOnEvent) string

var jobGerritTriggerOnEventsCodeFuncs = map[string]jobGerritTriggerOnEventsCodeFunc{}
var jobGerritTriggerOnEventsImportScriptFuncs = map[string]jobGerritTriggerOnEventsImportScriptFunc{}

func init() {
	ensureJobTriggerFuncs["*client.JobGerritTrigger"] = ensureJobGerritTrigger
}

func ensureJobGerritTriggerOnEvents(events *client.JobGerritTriggerOnEvents) error {
	if events == nil || events.Items == nil {
		return nil
	}
	for _, item := range *events.Items {
		if item.GetId() == "" {
			id, err := uuid.NewRandom()
			if err != nil {
				return err
			}
			item.SetId(id.String())
		}
	}
	return nil
}

func jobGerritTriggerOnEventsCode(
	propertyIndex int, triggerIndex int, events *client.JobGerritTriggerOnEvents,
) string {
	code := ""
	for _, item := range *events.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobGerritTriggerOnEventsCodeFuncs[reflectType]; ok {
			code += codeFunc(propertyIndex, triggerIndex, item)
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
