package client

import (
	"encoding/xml"
	"errors"
)

func init() {
	propertyUnmarshalFunc["org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty"] = unmarshalPipelineTriggersProperty
}

// ErrJobTriggerNotFound job property not found
var ErrJobTriggerNotFound = errors.New("Could not find job pipeline trigger")

type JobPipelineTriggersProperty struct {
	XMLName  xml.Name     `xml:"org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty"`
	Id       string       `xml:"id,attr,omitempty"`
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

func (p *JobPipelineTriggersProperty) SetId(id string) {
	p.Id = id
}

func unmarshalPipelineTriggersProperty(d *xml.Decoder, start xml.StartElement) (JobProperty, error) {
	property := NewJobPipelineTriggersProperty()
	err := d.DecodeElement(property, &start)
	if err != nil {
		return nil, err
	}
	return property, nil
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
