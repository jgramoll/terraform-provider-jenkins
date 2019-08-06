package client

import "encoding/xml"

func init() {
	jobGerritTriggerOnEventsUnmarshalFunc["com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginChangeMergedEvent"] = unmarshalJobGerritTriggerPluginChangeMergedEvent
}

type JobGerritTriggerPluginChangeMergedEvent struct {
	XMLName xml.Name `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginChangeMergedEvent"`
	Id      string   `xml:"id,attr,omitempty"`
}

func NewJobGerritTriggerPluginChangeMergedEvent() *JobGerritTriggerPluginChangeMergedEvent {
	return &JobGerritTriggerPluginChangeMergedEvent{}
}

func (event *JobGerritTriggerPluginChangeMergedEvent) GetId() string {
	return event.Id
}

func (e *JobGerritTriggerPluginChangeMergedEvent) SetId(id string) {
	e.Id = id
}

func unmarshalJobGerritTriggerPluginChangeMergedEvent(d *xml.Decoder, start xml.StartElement) (JobGerritTriggerOnEvent, error) {
	event := NewJobGerritTriggerPluginChangeMergedEvent()
	err := d.DecodeElement(event, &start)
	if err != nil {
		return nil, err
	}
	return event, nil
}
