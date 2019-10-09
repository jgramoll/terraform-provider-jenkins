package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobGerritTriggerOnEventsCodeFuncs["*client.JobGerritTriggerPluginDraftPublishedEvent"] = jobGerritTriggerPluginDraftPublishedEventCode
}

func jobGerritTriggerPluginDraftPublishedEventCode(
	triggerIndex string, e client.JobGerritTriggerOnEvent,
) string {
	return fmt.Sprintf(`
resource "jenkins_job_gerrit_trigger_draft_published_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_%v.id}"
}
`, triggerIndex)
}
