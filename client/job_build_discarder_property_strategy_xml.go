package client

import (
	"encoding/xml"
	"log"
)

type jobBuildDiscarderPropertyStrategyUnmarshaler func(*xml.Decoder, xml.StartElement) (JobBuildDiscarderPropertyStrategy, error)

var jobBuildDiscarderPropertyStrategyUnmarshalFunc = map[string]jobBuildDiscarderPropertyStrategyUnmarshaler{}

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
			if unmarshalXML, ok := jobBuildDiscarderPropertyStrategyUnmarshalFunc[attr.Value]; ok {
				newStrategyItem, err := unmarshalXML(d, start)
				if err != nil {
					return err
				}
				strategy.Item = newStrategyItem
				return nil
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
