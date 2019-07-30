package client

import "encoding/xml"

type JobGerritTriggerOnEvents struct {
	XMLName    xml.Name `xml:"triggerOnEvents"`
	LinkedList string   `xml:"class,attr"`
	Items      *[]JobGerritTriggerOnEvent
}

func NewJobGerritTriggerOnEvents() *JobGerritTriggerOnEvents {
	return &JobGerritTriggerOnEvents{
		LinkedList: "linked-list",
		Items:      &[]JobGerritTriggerOnEvent{},
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
	events.Items = &[]JobGerritTriggerOnEvent{}
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			// TODO use map
			switch elem.Name.Local {
			case "com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginDraftPublishedEvent":
				event := NewJobGerritTriggerPluginDraftPublishedEvent()
				err := d.DecodeElement(event, &elem)
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
