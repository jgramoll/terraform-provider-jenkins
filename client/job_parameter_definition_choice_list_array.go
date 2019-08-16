package client

import "encoding/xml"

type JobParameterDefinitionChoiceListArray struct {
	XMLName xml.Name  `xml:"a"`
	Class   string    `xml:"class,attr,omitempty"`
	Items   *[]string `xml:"string"`
}

func NewJobParameterDefinitionChoiceListArray() *JobParameterDefinitionChoiceListArray {
	return &JobParameterDefinitionChoiceListArray{
		Class: "string-array",
		Items: &[]string{},
	}
}

func (choices *JobParameterDefinitionChoiceListArray) Append(choice string) *JobParameterDefinitionChoiceListArray {
	newChoices := NewJobParameterDefinitionChoiceListArray()
	if choices != nil {
		*newChoices.Items = append(*choices.Items, choice)
	} else {
		*newChoices.Items = []string{choice}
	}
	return newChoices
}
