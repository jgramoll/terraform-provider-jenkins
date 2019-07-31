package client

import "encoding/xml"

type JobBuildDiscarderProperty struct {
	XMLName xml.Name `xml:"jenkins.model.BuildDiscarderProperty"`
	Id      string   `xml:"id,attr"`

	Strategy *JobBuildDiscarderPropertyStrategyXml `xml:"strategy"`
}

func NewJobBuildDiscarderProperty() *JobBuildDiscarderProperty {
	return &JobBuildDiscarderProperty{
		Strategy: NewJobBuildDiscarderPropertyStrategyXml(),
	}
}

func (property *JobBuildDiscarderProperty) GetId() string {
	return property.Id
}
