package client

import (
	"encoding/xml"
	"errors"
)

// ErrJobTriggerNotFound job property not found
var ErrJobTriggerNotFound = errors.New("Could not find job pipeline trigger")

type JobPipelineTriggersProperty struct {
	XMLName  xml.Name     `xml:"org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty"`
	Id       string       `xml:"id,attr"`
	Triggers *JobTriggers `xml:"triggers"`
}

func NewJobPipelineTriggersProperty() *JobPipelineTriggersProperty {
	return &JobPipelineTriggersProperty{
		Triggers: NewJobTriggers(),
	}
}

func (property *JobPipelineTriggersProperty) GetId() string {
	return property.Id
}

func (property *JobPipelineTriggersProperty) GetTrigger(triggerId string) (JobTrigger, error) {
	triggers := *(property.Triggers).Items
	for _, trigger := range triggers {
		if trigger.GetId() == triggerId {
			return trigger, nil
		}
	}
	return nil, ErrJobTriggerNotFound
}

func (property *JobPipelineTriggersProperty) UpdateTrigger(newTrigger JobTrigger) error {
	triggerId := newTrigger.GetId()
	triggers := *(property.Triggers).Items
	for i, trigger := range triggers {
		if trigger.GetId() == triggerId {
			triggers[i] = newTrigger
			return nil
		}
	}
	return ErrJobTriggerNotFound
}

func (property *JobPipelineTriggersProperty) DeleteTrigger(triggerId string) error {
	triggers := *(property.Triggers).Items
	for i, trigger := range triggers {
		if trigger.GetId() == triggerId {
			*property.Triggers.Items = append(triggers[:i], triggers[i+1:]...)
			return nil
		}
	}
	return ErrJobTriggerNotFound
}
