package client

import "encoding/xml"

type JobPipelineBuildDiscarderProperty struct {
	XMLName xml.Name `xml:"jenkins.model.BuildDiscarderProperty"`
	Id      string   `xml:"id,attr"`

	Strategy JobPipelineBuildDiscarderPropertyStrategy `xml:"strategy"`
}

func NewJobPipelineBuildDiscarderProperty() *JobPipelineBuildDiscarderProperty {
	return &JobPipelineBuildDiscarderProperty{}
}

func (property *JobPipelineBuildDiscarderProperty) GetId() string {
	return property.Id
}
