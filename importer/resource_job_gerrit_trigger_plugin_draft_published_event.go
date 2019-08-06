package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	jobGerritTriggerOnEventsCodeFuncs["*client.JobGerritTriggerPluginDraftPublishedEvent"] = jobGerritTriggerPluginDraftPublishedEventCode
	jobGerritTriggerOnEventsImportScriptFuncs["*client.JobGerritTriggerPluginDraftPublishedEvent"] = jobGerritTriggerPluginDraftPublishedEventImportScript
}

func jobGerritTriggerPluginDraftPublishedEventCode(client.JobGerritTriggerOnEvent) string {
	return `
resource "jenkins_job_gerrit_trigger_draft_published_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.main.id}"
}
`
}

func jobGerritTriggerPluginDraftPublishedEventImportScript(
	jobName string, propertyId string, triggerId string,
	e client.JobGerritTriggerOnEvent,
) string {
	return fmt.Sprintf(`
terraform import jenkins_job_gerrit_trigger_draft_published_event.main "%v"
`, provider.ResourceJobGerritTriggerEventId(jobName, propertyId, triggerId, e.GetId()))
}
