package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobGerritTriggerOnEventsCodeFuncs["*client.JobGerritTriggerPluginDraftPublishedEvent"] = jobGerritTriggerPluginDraftPublishedEventCode
}

func jobGerritTriggerPluginDraftPublishedEventCode(e client.JobGerritTriggerOnEvent) string {
	return `
      trigger_on_event {
        type = "PluginDraftPublishedEvent"
      }
`
}
