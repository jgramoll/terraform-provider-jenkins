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
	var newEventItems []JobGerritTriggerOnEvent
	if events.Items != nil {
		newEventItems = append(*events.Items, event)
	} else {
		newEventItems = append(newEventItems, event)
	}
	newEvents := NewJobGerritTriggerOnEvents()
	newEvents.Items = &newEventItems
	return newEvents
}
