package client

import "encoding/xml"

type JobGerritTriggerOnEvents struct {
	XMLName    xml.Name `xml:"triggerOnEvents"`
	LinkedList string   `xml:"class,attr"`
}

func NewJobGerritTriggerOnEvents() *JobGerritTriggerOnEvents {
	return &JobGerritTriggerOnEvents{
		LinkedList: "linked-list",
	}
}
