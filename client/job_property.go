package client

import (
	"encoding/xml"
)

type JobProperties struct {
	XMLName xml.Name       `xml:"properties"`
	Items   *[]JobProperty `xml:",any"`
}
type JobProperty interface {
	GetId() string
}

func NewJobProperties() *JobProperties {
	return &JobProperties{
		Items: &[]JobProperty{},
	}
}

func (properties *JobProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tok xml.Token
	var err error
	properties.Items = &[]JobProperty{}
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			var property JobProperty
			// TODO use map
			switch elem.Name.Local {
			default:
				continue
			case "org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty":
				property = &JobPipelineTriggersProperty{}
			}
			d.DecodeElement(property, &elem)
			*properties.Items = append(*(*properties).Items, property)
		}
		if end, ok := tok.(xml.EndElement); ok {
			if end.Name.Local == "properties" {
				break
			}
		}
	}
	return err
}
