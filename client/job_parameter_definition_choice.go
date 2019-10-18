package client

import "encoding/xml"

func init() {
	parametersDefinitionUnmarshalFunc["hudson.model.ChoiceParameterDefinition"] = unmarshalJobParameterDefinitionChoice
}

type JobParameterDefinitionChoice struct {
	XMLName xml.Name `xml:"hudson.model.ChoiceParameterDefinition"`

	Name        string `xml:"name"`
	Description string `xml:"description"`

	Choices *JobParameterDefinitionChoiceList `xml:"choices"`
}

func NewJobParameterDefinitionChoice() *JobParameterDefinitionChoice {
	return &JobParameterDefinitionChoice{
		Choices: NewJobParameterDefinitionChoiceList(),
	}
}

func (d *JobParameterDefinitionChoice) GetType() JobParameterDefinitionType {
	return ChoiceParameterDefinitionType
}

func unmarshalJobParameterDefinitionChoice(d *xml.Decoder, start xml.StartElement) (JobParameterDefinition, error) {
	definition := NewJobParameterDefinitionChoice()
	err := d.DecodeElement(definition, &start)
	if err != nil {
		return nil, err
	}
	return definition, nil
}
