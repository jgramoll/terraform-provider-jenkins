package client

import (
	"encoding/xml"
	"log"
)

type JobDefinitionXml struct {
	Item JobDefinition
}

type JobDefinition interface{}

func (jobDefinition *JobDefinitionXml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(jobDefinition.Item, start)
}

func (jobDefinition *JobDefinitionXml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "class":
			switch attr.Value {
			default:
			case "org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition":
				newDef := NewCpsScmFlowDefinition()
				jobDefinition.Item = newDef
				return d.DecodeElement(newDef, &start)
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
