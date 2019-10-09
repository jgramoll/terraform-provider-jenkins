package client

import (
	"encoding/xml"
	"errors"
)

func init() {
	propertyUnmarshalFunc["hudson.model.ParametersDefinitionProperty"] = unmarshalJobParametersDefinitionProperty
}

// ErrJobParameterDefinitionNotFound job parameter definition not found
var ErrJobParameterDefinitionNotFound = errors.New("Could not find job parameter definition")

type JobParametersDefinitionProperty struct {
	XMLName xml.Name `xml:"hudson.model.ParametersDefinitionProperty"`
	Plugin  string   `xml:"plugin,attr,omitempty"`

	ParameterDefinitions *JobParameterDefinitions `xml:"parameterDefinitions"`
}

func NewJobParametersDefinitionProperty() *JobParametersDefinitionProperty {
	return &JobParametersDefinitionProperty{}
}

func (property *JobParametersDefinitionProperty) GetType() JobPropertyType {
	return ParametersDefinitionPropertyType
}

func unmarshalJobParametersDefinitionProperty(d *xml.Decoder, start xml.StartElement) (JobProperty, error) {
	property := NewJobParametersDefinitionProperty()
	err := d.DecodeElement(property, &start)
	if err != nil {
		return nil, err
	}
	return property, nil
}
