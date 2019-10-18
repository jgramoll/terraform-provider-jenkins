package client

import "encoding/xml"

func init() {
	jobGerritTriggerOnEventsUnmarshalFunc["com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginDraftPublishedEvent"] = unmarshalJobGerritTriggerPluginDraftPublishedEvent
}

type JobGerritTriggerPluginDraftPublishedEvent struct {
	XMLName xml.Name `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginDraftPublishedEvent"`
}

func NewJobGerritTriggerPluginDraftPublishedEvent() *JobGerritTriggerPluginDraftPublishedEvent {
	return &JobGerritTriggerPluginDraftPublishedEvent{}
}

func (event *JobGerritTriggerPluginDraftPublishedEvent) GetType() JobGerritTriggerOnEventType {
	return PluginDraftPublishedEventType
}

func unmarshalJobGerritTriggerPluginDraftPublishedEvent(d *xml.Decoder, start xml.StartElement) (JobGerritTriggerOnEvent, error) {
	event := NewJobGerritTriggerPluginDraftPublishedEvent()
	err := d.DecodeElement(event, &start)
	if err != nil {
		return nil, err
	}
	return event, nil
}
