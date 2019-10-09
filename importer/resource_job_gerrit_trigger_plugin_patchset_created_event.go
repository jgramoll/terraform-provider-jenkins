package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobGerritTriggerOnEventsCodeFuncs["*client.JobGerritTriggerPluginPatchsetCreatedEvent"] = jobGerritTriggerPluginPatchsetCreatedEventCode
}

func jobGerritTriggerPluginPatchsetCreatedEventCode(
	triggerIndex string, eventInterface client.JobGerritTriggerOnEvent,
) string {
	event := eventInterface.(*client.JobGerritTriggerPluginPatchsetCreatedEvent)
	return fmt.Sprintf(`
resource "jenkins_job_gerrit_trigger_patchset_created_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_%v.id}"

	exclude_drafts         = %v
	exclude_trivial_rebase = %v
	exclude_no_code_change = %v
	exclude_private_state  = %v
	exclude_wip_state      = %v
}
`, triggerIndex,
		event.ExcludeDrafts, event.ExcludeTrivialRebase, event.ExcludeNoCodeChange,
		event.ExcludePrivateState, event.ExcludeWipState)
}
