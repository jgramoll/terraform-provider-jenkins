package client

import (
	"encoding/xml"
)

type JobProperties struct {
	XMLName xml.Name       `xml:"properties"`
	Items   *[]JobProperty `xml:",any"`
}

func NewJobProperties() *JobProperties {
	return &JobProperties{
		Items: &[]JobProperty{},
	}
}

func (properties *JobProperties) Append(property JobProperty) *JobProperties {
	newProperties := NewJobProperties()
	if properties.Items != nil {
		*newProperties.Items = append(*properties.Items, property)
	} else {
		*newProperties.Items = []JobProperty{property}
	}
	return newProperties
}

func (properties *JobProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tok xml.Token
	var err error
	properties.Items = &[]JobProperty{}
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			// TODO use map
			switch elem.Name.Local {
			case "jenkins.model.BuildDiscarderProperty":
				property := NewJobBuildDiscarderProperty()
				err := d.DecodeElement(property, &elem)
				if err != nil {
					return err
				}
				*properties.Items = append(*(properties).Items, property)
			case "org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty":
				property := NewJobPipelineTriggersProperty()
				err := d.DecodeElement(property, &elem)
				if err != nil {
					return err
				}
				*properties.Items = append(*(properties).Items, property)
			}
		}
		if end, ok := tok.(xml.EndElement); ok {
			if end.Name.Local == "properties" {
				break
			}
		}
	}
	return err
}
