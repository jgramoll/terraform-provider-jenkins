package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	ensureJobPropertyFuncs["*client.JobPipelineTriggersProperty"] = ensureJobPipelineTriggersProperty
	jobPropertyCodeFuncs["*client.JobPipelineTriggersProperty"] = jobPipelineTriggersPropertyCode
	jobPropertyImportScriptFuncs["*client.JobPipelineTriggersProperty"] = jobPipelineTriggersPropertyImportScript
}

type ensureJobTriggerFunc func(client.JobTrigger) error
type jobTriggerCodeFunc func(int, int, client.JobTrigger) string
type jobTriggerImportScriptFunc func(int, int, string, string, client.JobTrigger) string

var ensureJobTriggerFuncs = map[string]ensureJobTriggerFunc{}
var jobTriggerCodeFuncs = map[string]jobTriggerCodeFunc{}
var jobTriggerImportScriptFuncs = map[string]jobTriggerImportScriptFunc{}

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

func jobPipelineTriggersPropertyCode(propertyIndex int, propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobPipelineTriggersProperty)

	triggersCode := ""
	for i, trigger := range *property.Triggers.Items {
		reflectType := reflect.TypeOf(trigger).String()
		if triggerCodeFunc, ok := jobTriggerCodeFuncs[reflectType]; ok {
			triggersCode += triggerCodeFunc(propertyIndex, i+1, trigger)
		} else {
			log.Println("[WARNING] Unknown Job Trigger Type", reflectType)
		}
	}
	return fmt.Sprintf(`
resource "jenkins_job_pipeline_triggers_property" "property_%v" {
	job = "${jenkins_job.main.name}"
}
`, propertyIndex) + triggersCode
}

func jobPipelineTriggersPropertyImportScript(propertyIndex int, jobName string, propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobPipelineTriggersProperty)

	triggersCode := ""
	for i, trigger := range *property.Triggers.Items {
		reflectType := reflect.TypeOf(trigger).String()
		if triggerCodeFunc, ok := jobTriggerImportScriptFuncs[reflectType]; ok {
			triggersCode += triggerCodeFunc(propertyIndex, i+1, jobName, property.Id, trigger)
		} else {
			log.Println("[WARNING] Unknown Job Trigger Type", reflectType)
		}
	}
	return fmt.Sprintf(`
terraform import jenkins_job_pipeline_triggers_property.property_%v "%v"
`, propertyIndex, provider.ResourceJobPropertyId(jobName, property.Id)) + triggersCode
}
