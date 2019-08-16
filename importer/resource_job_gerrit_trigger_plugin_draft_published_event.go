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

func jobGerritTriggerPluginDraftPublishedEventCode(
	propertyIndex int, triggerIndex int, e client.JobGerritTriggerOnEvent,
) string {
	return fmt.Sprintf(`
resource "jenkins_job_gerrit_trigger_draft_published_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_%v_%v.id}"
}
`, propertyIndex, triggerIndex)
}

func jobGerritTriggerPluginDraftPublishedEventImportScript(
	jobName string, propertyId string, triggerId string,
	e client.JobGerritTriggerOnEvent,
) string {
	return fmt.Sprintf(`
terraform import jenkins_job_gerrit_trigger_draft_published_event.main "%v"
`, provider.ResourceJobGerritTriggerEventId(jobName, propertyId, triggerId, e.GetId()))
}
