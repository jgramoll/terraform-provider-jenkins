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
	Triggers *JobTriggers `xml:"triggers"`
}

func NewJobPipelineTriggersProperty() *JobPipelineTriggersProperty {
	return &JobPipelineTriggersProperty{
		Triggers: NewJobTriggers(),
	}
}

func (property *JobPipelineTriggersProperty) GetType() JobPropertyType {
	return PipelineTriggersJobPropertyType
}

func unmarshalPipelineTriggersProperty(d *xml.Decoder, start xml.StartElement) (JobProperty, error) {
	property := NewJobPipelineTriggersProperty()
	err := d.DecodeElement(property, &start)
	if err != nil {
		return nil, err
	}
	return property, nil
}
