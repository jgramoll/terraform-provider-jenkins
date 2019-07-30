package client

import "encoding/xml"

type JobGerritTriggerPluginDraftPublishedEvent struct {
	XMLName xml.Name `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginDraftPublishedEvent"`
	Id      string   `xml:"id,attr"`
}

func NewJobGerritTriggerPluginDraftPublishedEvent() *JobGerritTriggerPluginDraftPublishedEvent {
	return &JobGerritTriggerPluginDraftPublishedEvent{}
}

func (event *JobGerritTriggerPluginDraftPublishedEvent) GetId() string {
	return event.Id
}
