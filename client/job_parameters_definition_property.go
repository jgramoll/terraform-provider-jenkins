package client

import (
	"encoding/xml"
)

func init() {
	propertyUnmarshalFunc["hudson.model.ParametersDefinitionProperty"] = unmarshalJobParametersDefinitionProperty
}

type JobParametersDefinitionProperty struct {
	XMLName xml.Name `xml:"hudson.model.ParametersDefinitionProperty"`
	Id      string   `xml:"id,attr,omitempty"`
	Plugin  string   `xml:"plugin,attr,omitempty"`

	ParameterDefinitions *JobParameterDefinitions `xml:"parameterDefinitions"`
}

func NewJobParametersDefinitionProperty() *JobParametersDefinitionProperty {
	return &JobParametersDefinitionProperty{}
}

func (property *JobParametersDefinitionProperty) GetId() string {
	return property.Id
}

func (p *JobParametersDefinitionProperty) SetId(id string) {
	p.Id = id
}

func unmarshalJobParametersDefinitionProperty(d *xml.Decoder, start xml.StartElement) (JobProperty, error) {
	property := NewJobParametersDefinitionProperty()
	err := d.DecodeElement(property, &start)
	if err != nil {
		return nil, err
	}
	return property, nil
}
