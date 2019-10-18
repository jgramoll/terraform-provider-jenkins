package main

import (
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerOnEventsCodeFunc func(client.JobGerritTriggerOnEvent) string

var jobGerritTriggerOnEventsCodeFuncs = map[string]jobGerritTriggerOnEventsCodeFunc{}

func jobGerritTriggerOnEventsCode(events *client.JobGerritTriggerOnEvents) string {
	code := ""
	for _, item := range *events.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobGerritTriggerOnEventsCodeFuncs[reflectType]; ok {
			code += codeFunc(item)
		} else {
			log.Println("[WARNING] Unknown gerrit trigger on event type:", reflectType)
		}
	}
	return code
}
