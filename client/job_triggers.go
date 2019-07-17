package client

import (
	"encoding/xml"
)

type JobTriggers struct {
	XMLName xml.Name `xml:"triggers"`
	Items   *[]JobTrigger
}

func NewJobTriggers() *JobTriggers {
	return &JobTriggers{
		Items: &[]JobTrigger{},
	}
}

func (triggers *JobTriggers) Append(trigger JobTrigger) *JobTriggers {
	var newTriggerItems []JobTrigger
	if triggers.Items != nil {
		newTriggerItems = append(*triggers.Items, trigger)
	} else {
		newTriggerItems = append(newTriggerItems, trigger)
	}
	newTriggers := NewJobTriggers()
	newTriggers.Items = &newTriggerItems
	return newTriggers
}

func (triggers *JobTriggers) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tok xml.Token
	var err error
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			// TODO use map
			switch elem.Name.Local {
			case "com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.GerritTrigger":
				trigger := NewJobGerritTrigger()
				err := d.DecodeElement(trigger, &elem)
				if err != nil {
					return err
				}
				*triggers.Items = append(*(*triggers).Items, trigger)
			}
		}
		if end, ok := tok.(xml.EndElement); ok {
			if end.Name.Local == "triggers" {
				break
			}
		}
	}
	return err
}
