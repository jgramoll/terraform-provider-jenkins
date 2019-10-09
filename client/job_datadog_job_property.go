package client

import (
	"encoding/xml"
)

func init() {
	propertyUnmarshalFunc["org.datadog.jenkins.plugins.datadog.DatadogJobProperty"] = unmarshalJobDatadogJobProperty
}

type JobDatadogJobProperty struct {
	XMLName xml.Name `xml:"org.datadog.jenkins.plugins.datadog.DatadogJobProperty"`
	Plugin  string   `xml:"plugin,attr,omitempty"`

	EmitOnCheckout bool `xml:"emitOnCheckout"`
}

func NewJobDatadogJobProperty() *JobDatadogJobProperty {
	return &JobDatadogJobProperty{
		EmitOnCheckout: false,
	}
}

func (property *JobDatadogJobProperty) GetType() JobPropertyType {
	return DatadogJobPropertyType
}

func unmarshalJobDatadogJobProperty(d *xml.Decoder, start xml.StartElement) (JobProperty, error) {
	property := NewJobDatadogJobProperty()
	err := d.DecodeElement(property, &start)
	if err != nil {
		return nil, err
	}
	return property, nil
}
