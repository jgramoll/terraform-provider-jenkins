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

func (property *JobParametersDefinitionProperty) GetParameterDefinition(
	definitionId string,
) (JobParameterDefinition, error) {
	definitions := *(property.ParameterDefinitions).Items
	for _, definition := range definitions {
		if definition.GetId() == definitionId {
			return definition, nil
		}
	}
	return nil, ErrJobParameterDefinitionNotFound
}

func (property *JobParametersDefinitionProperty) UpdateParameterDefinition(
	newDefinition JobParameterDefinition,
) error {
	definitionId := newDefinition.GetId()
	definitions := *(property.ParameterDefinitions).Items
	for i, definition := range definitions {
		if definition.GetId() == definitionId {
			definitions[i] = newDefinition
			return nil
		}
	}
	return ErrJobParameterDefinitionNotFound
}

func (property *JobParametersDefinitionProperty) DeleteParameterDefinition(definitionId string) error {
	definitions := *(property.ParameterDefinitions).Items
	for i, definition := range definitions {
		if definition.GetId() == definitionId {
			*property.ParameterDefinitions.Items = append(definitions[:i], definitions[i+1:]...)
			return nil
		}
	}
	return ErrJobParameterDefinitionNotFound
}
