package client

import (
	"encoding/xml"
)

type jobTriggerUnmarshaler func(*xml.Decoder, xml.StartElement) (JobTrigger, error)

var jobTriggerUnmarshalFunc = map[string]jobTriggerUnmarshaler{}

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
	newTriggers := NewJobTriggers()
	if triggers.Items != nil {
		*newTriggers.Items = append(*triggers.Items, trigger)
	} else {
		*newTriggers.Items = append([]JobTrigger{}, trigger)
	}
	return newTriggers
}

func (triggers *JobTriggers) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tok xml.Token
	var err error
	*triggers = *NewJobTriggers()
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			if unmarshalXML, ok := jobTriggerUnmarshalFunc[elem.Name.Local]; ok {
				trigger, err := unmarshalXML(d, elem)
				if err != nil {
					return err
				}
				*triggers.Items = append(*triggers.Items, trigger)
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
