package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	jobGerritTriggerOnEventsCodeFuncs["*client.JobGerritTriggerPluginPatchsetCreatedEvent"] = jobGerritTriggerPluginPatchsetCreatedEventCode
	jobGerritTriggerOnEventsImportScriptFuncs["*client.JobGerritTriggerPluginPatchsetCreatedEvent"] = jobGerritTriggerPluginPatchsetCreatedEventImportScript
}

func jobGerritTriggerPluginPatchsetCreatedEventCode(eventInterface client.JobGerritTriggerOnEvent) string {
	event := eventInterface.(*client.JobGerritTriggerPluginPatchsetCreatedEvent)
	return fmt.Sprintf(`
resource "jenkins_job_gerrit_trigger_patchset_created_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.main.id}"

	exclude_drafts         = %v
	exclude_trivial_rebase = %v
	exclude_no_code_change = %v
	exclude_private_state  = %v
	exclude_wip_state      = %v
}
`, event.ExcludeDrafts, event.ExcludeTrivialRebase, event.ExcludeNoCodeChange,
		event.ExcludePrivateState, event.ExcludeWipState)
}
func jobGerritTriggerPluginPatchsetCreatedEventImportScript(
	jobName string, propertyId string, triggerId string,
	eventInterface client.JobGerritTriggerOnEvent,
) string {
	event := eventInterface.(*client.JobGerritTriggerPluginPatchsetCreatedEvent)
	return fmt.Sprintf(`
terraform import jenkins_job_gerrit_trigger_patchset_created_event.main "%v"
`, provider.ResourceJobGerritTriggerEventId(jobName, propertyId, triggerId, event.Id))
}
