package client

import "encoding/xml"

type JobGerritTriggerPluginDraftPublishedEvent struct {
	XMLName xml.Name `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginDraftPublishedEvent"`
}

func NewJobGerritTriggerPluginDraftPublishedEvent() *JobGerritTriggerPluginDraftPublishedEvent {
	return &JobGerritTriggerPluginDraftPublishedEvent{}
}
