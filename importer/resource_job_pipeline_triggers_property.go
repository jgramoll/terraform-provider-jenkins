package main

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	ensureJobPropertyFuncs["*client.JobPipelineTriggersProperty"] = ensureJobPipelineTriggersProperty
}

type ensureJobTriggerFunc func(client.JobTrigger) error

var ensureJobTriggerFuncs = map[string]ensureJobTriggerFunc{}

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
