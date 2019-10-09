package client

import (
	"encoding/xml"
)

func init() {
	propertyUnmarshalFunc["org.jenkinsci.plugins.workflow.job.properties.DisableConcurrentBuildsJobProperty"] = unmarshalJobDisableConcurrentBuildsJobProperty
}

type JobDisableConcurrentBuildsJobProperty struct {
	XMLName xml.Name `xml:"org.jenkinsci.plugins.workflow.job.properties.DisableConcurrentBuildsJobProperty"`
	Plugin  string   `xml:"plugin,attr,omitempty"`
}

func NewJobDisableConcurrentBuildsJobProperty() *JobDisableConcurrentBuildsJobProperty {
	return &JobDisableConcurrentBuildsJobProperty{}
}

func (property *JobDisableConcurrentBuildsJobProperty) GetType() JobPropertyType {
	return DisableConcurrentBuildsJobPropertyType
}

func unmarshalJobDisableConcurrentBuildsJobProperty(d *xml.Decoder, start xml.StartElement) (JobProperty, error) {
	property := NewJobDisableConcurrentBuildsJobProperty()
	err := d.DecodeElement(property, &start)
	if err != nil {
		return nil, err
	}
	return property, nil
}
