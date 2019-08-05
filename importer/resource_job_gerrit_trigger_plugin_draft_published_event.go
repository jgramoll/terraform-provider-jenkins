package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobGerritTriggerOnEventsCodeFuncs["*client.JobGerritTriggerPluginDraftPublishedEvent"] = jobGerritTriggerPluginDraftPublishedEventCode
}

func jobGerritTriggerPluginDraftPublishedEventCode(client.JobGerritTriggerOnEvent) string {
	return `
resource "jenkins_job_gerrit_trigger_draft_published_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.main.id}"
}
`
}
