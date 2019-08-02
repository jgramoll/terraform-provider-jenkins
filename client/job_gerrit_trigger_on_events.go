package client

import "encoding/xml"

type jobGerritTriggerOnEventsUnmarshaler func(*xml.Decoder, xml.StartElement) (JobGerritTriggerOnEvent, error)

var jobGerritTriggerOnEventsUnmarshalFunc = map[string]jobGerritTriggerOnEventsUnmarshaler{}

type JobGerritTriggerOnEvents struct {
	XMLName xml.Name `xml:"triggerOnEvents"`
	Class   string   `xml:"class,attr"`
	Items   *[]JobGerritTriggerOnEvent
}

func NewJobGerritTriggerOnEvents() *JobGerritTriggerOnEvents {
	return &JobGerritTriggerOnEvents{
		Class: "linked-list",
		Items: &[]JobGerritTriggerOnEvent{},
	}
}

func (events *JobGerritTriggerOnEvents) Append(event JobGerritTriggerOnEvent) *JobGerritTriggerOnEvents {
	newEvents := NewJobGerritTriggerOnEvents()
	if events.Items != nil {
		*newEvents.Items = append(*events.Items, event)
	} else {
		*newEvents.Items = []JobGerritTriggerOnEvent{event}
	}
	return newEvents
}

func (events *JobGerritTriggerOnEvents) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tok xml.Token
	var err error
	*events = *NewJobGerritTriggerOnEvents()
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			if unmarshalXML, ok := jobGerritTriggerOnEventsUnmarshalFunc[elem.Name.Local]; ok {
				event, err := unmarshalXML(d, elem)
				if err != nil {
					return err
				}
				*events.Items = append(*events.Items, event)
			}
		}
		if end, ok := tok.(xml.EndElement); ok {
			if end.Name.Local == "triggerOnEvents" {
				break
			}
		}
	}
	return err
}
