package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	ensureJobPropertyFuncs["*client.JobPipelineTriggersProperty"] = ensureJobPipelineTriggersProperty
	jobPropertyCodeFuncs["*client.JobPipelineTriggersProperty"] = jobPipelineTriggersPropertyCode
}

type ensureJobTriggerFunc func(client.JobTrigger) error
type jobTriggerCodeFunc func(client.JobTrigger) string

var ensureJobTriggerFuncs = map[string]ensureJobTriggerFunc{}
var jobTriggerCodeFuncs = map[string]jobTriggerCodeFunc{}

func ensureJobPipelineTriggersProperty(propertyInterface client.JobProperty) error {
	property := propertyInterface.(*client.JobPipelineTriggersProperty)
	for _, trigger := range *property.Triggers.Items {
		reflectType := reflect.TypeOf(trigger).String()
		ensureFunc, ok := ensureJobTriggerFuncs[reflectType]
		if !ok {
			return fmt.Errorf("Unknown Job Trigger Type %s", reflectType)
		}
		if err := ensureFunc(trigger); err != nil {
			return err
		}
	}
	return nil
}

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
	return `
resource "jenkins_job_pipeline_triggers_property" "main" {
	job = "${jenkins_job.main.id}"
}
` + triggersCode
}
