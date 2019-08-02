package client

import (
	"encoding/xml"
)

type propertyUnmarshaler func(*xml.Decoder, xml.StartElement) (JobProperty, error)

var propertyUnmarshalFunc = map[string]propertyUnmarshaler{}

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
	*properties = *NewJobProperties()
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			if unmarshalXML, ok := propertyUnmarshalFunc[elem.Name.Local]; ok {
				property, err := unmarshalXML(d, elem)
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
