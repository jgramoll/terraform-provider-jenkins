package client

import "encoding/xml"

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
