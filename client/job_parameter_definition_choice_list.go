package client

import "encoding/xml"

type JobParameterDefinitionChoiceList struct {
	XMLName xml.Name                               `xml:"choices"`
	Class   string                                 `xml:"class,attr,omitempty"`
	Items   *JobParameterDefinitionChoiceListArray `xml:"a"`
}

func NewJobParameterDefinitionChoiceList() *JobParameterDefinitionChoiceList {
	return &JobParameterDefinitionChoiceList{
		Class: "java.util.Arrays$ArrayList",
		Items: NewJobParameterDefinitionChoiceListArray(),
	}
}

func (choices *JobParameterDefinitionChoiceList) Append(choice string) *JobParameterDefinitionChoiceList {
	newChoices := NewJobParameterDefinitionChoiceList()
	if choices != nil {
		newChoices.Items = choices.Items.Append(choice)
	} else {
		newChoices.Items = NewJobParameterDefinitionChoiceListArray().Append(choice)
	}
	return newChoices
}
