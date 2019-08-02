package client

import (
	"encoding/xml"
	"log"
)

type jobDefinitionUnmarshaler func(*xml.Decoder, xml.StartElement) (JobDefinition, error)

var jobDefinitionUnmarshalFunc = map[string]jobDefinitionUnmarshaler{}

type JobDefinitionXml struct {
	Item JobDefinition
}

type JobDefinition interface {
	GetId() string
	SetId(string)
}

func (jobDefinition *JobDefinitionXml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(jobDefinition.Item, start)
}

func (jobDefinition *JobDefinitionXml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "class":
			if unmarshalXML, ok := jobDefinitionUnmarshalFunc[attr.Value]; ok {
				newDefItem, err := unmarshalXML(d, start)
				if err != nil {
					return err
				}
				jobDefinition.Item = newDefItem
				return nil
			}
		}
	}

	log.Println("[WARN] Unable to unmarshal definition xml. Unknown or missing class")
	var err error
	// must read all of doc or it complains
	for _, err := d.Token(); err == nil; _, err = d.Token() {
	}
	return err
}
