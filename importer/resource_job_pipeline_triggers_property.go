package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobPipelineTriggersProperty"] = jobPipelineTriggersPropertyCode
}

type jobTriggerCodeFunc func(client.JobTrigger) string

var jobTriggerCodeFuncs = map[string]jobTriggerCodeFunc{}

func jobPipelineTriggersPropertyCode(propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobPipelineTriggersProperty)

	triggersCode := ""
	for _, trigger := range *property.Triggers.Items {
		reflectType := reflect.TypeOf(trigger).String()
		if triggerCodeFunc, ok := jobTriggerCodeFuncs[reflectType]; ok {
			triggersCode += triggerCodeFunc(trigger)
		} else {
			log.Println("[WARNING] Unknown Job Trigger Type", reflectType)
		}
	}
	return fmt.Sprintf(`
  property {
    type = "PipelineTriggersJobProperty"
%s  }
`, triggersCode)
}
