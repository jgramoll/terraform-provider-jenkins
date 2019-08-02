package client

import (
	"encoding/xml"
)

func init() {
	propertyUnmarshalFunc["org.datadog.jenkins.plugins.datadog.DatadogJobProperty"] = unmarshalJobDatadogJobProperty
}

type JobDatadogJobProperty struct {
	XMLName xml.Name `xml:"org.datadog.jenkins.plugins.datadog.DatadogJobProperty"`
	Id      string   `xml:"id,attr,omitempty"`
	Plugin  string   `xml:"plugin,attr,omitempty"`

	EmitOnCheckout bool `xml:"emitOnCheckout"`
}

func NewJobDatadogJobProperty() *JobDatadogJobProperty {
	return &JobDatadogJobProperty{
		EmitOnCheckout: false,
	}
}

func (property *JobDatadogJobProperty) GetId() string {
	return property.Id
}

func (p *JobDatadogJobProperty) SetId(id string) {
	p.Id = id
}

func unmarshalJobDatadogJobProperty(d *xml.Decoder, start xml.StartElement) (JobProperty, error) {
	property := NewJobDatadogJobProperty()
	err := d.DecodeElement(property, &start)
	if err != nil {
		return nil, err
	}
	return property, nil
}
