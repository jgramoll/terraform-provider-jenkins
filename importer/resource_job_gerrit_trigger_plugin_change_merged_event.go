package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobGerritTriggerOnEventsCodeFuncs["*client.JobGerritTriggerPluginChangeMergedEvent"] = jobGerritTriggerPluginChangeMergedEventCode
}

func jobGerritTriggerPluginChangeMergedEventCode(e client.JobGerritTriggerOnEvent) string {
	return `
      trigger_on_event {
        type = "PluginChangeMergedEvent"
      }
`
}
