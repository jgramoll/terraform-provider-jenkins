package client

import (
	"encoding/xml"
	"log"
)

type JobBuildDiscarderPropertyStrategyXml struct {
	Item JobBuildDiscarderPropertyStrategy `xml:",any"`
}

func NewJobBuildDiscarderPropertyStrategyXml() *JobBuildDiscarderPropertyStrategyXml {
	return &JobBuildDiscarderPropertyStrategyXml{}
}

func (strategy *JobBuildDiscarderPropertyStrategyXml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(strategy.Item, start)
}

func (strategy *JobBuildDiscarderPropertyStrategyXml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "class":
			switch attr.Value {
			default:
			case "hudson.tasks.LogRotator":
				newStrategy := NewJobBuildDiscarderPropertyStrategyLogRotator()
				strategy.Item = newStrategy
				return d.DecodeElement(newStrategy, &start)
			}
		}
	}

	log.Println("[WARN] Unable to unmarshal build discarder property strategy xml. Unknown or missing class")
	var err error
	// must read all of doc or it complains
	for _, err := d.Token(); err == nil; _, err = d.Token() {
	}
	return err
}
