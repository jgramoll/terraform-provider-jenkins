package client

import "encoding/xml"

type parametersDefinitionUnmarshaler func(*xml.Decoder, xml.StartElement) (JobParameterDefinition, error)

var parametersDefinitionUnmarshalFunc = map[string]parametersDefinitionUnmarshaler{}

type JobParameterDefinitions struct {
	XMLName xml.Name `xml:"parameterDefinitions"`
	Items   *[]JobParameterDefinition
}

func NewJobParameterDefinitions() *JobParameterDefinitions {
	return &JobParameterDefinitions{
		Items: &[]JobParameterDefinition{},
	}
}

func (defs *JobParameterDefinitions) Append(def JobParameterDefinition) *JobParameterDefinitions {
	newDefs := NewJobParameterDefinitions()
	if defs != nil {
		*newDefs.Items = append(*defs.Items, def)
	} else {
		*newDefs.Items = []JobParameterDefinition{def}
	}
	return newDefs
}

func (defs *JobParameterDefinitions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tok xml.Token
	var err error
	*defs = *NewJobParameterDefinitions()
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			if unmarshalXML, ok := parametersDefinitionUnmarshalFunc[elem.Name.Local]; ok {
				def, err := unmarshalXML(d, elem)
				if err != nil {
					return err
				}
				*defs.Items = append(*(defs).Items, def)
			}
		}
		if end, ok := tok.(xml.EndElement); ok {
			if end.Name.Local == "parameterDefinitions" {
				break
			}
		}
	}
	return err
}
