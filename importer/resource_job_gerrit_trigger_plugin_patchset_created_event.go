package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobGerritTriggerOnEventsCodeFuncs["*client.JobGerritTriggerPluginPatchsetCreatedEvent"] = jobGerritTriggerPluginPatchsetCreatedEventCode
}

func jobGerritTriggerPluginPatchsetCreatedEventCode(eventInterface client.JobGerritTriggerOnEvent) string {
	event := eventInterface.(*client.JobGerritTriggerPluginPatchsetCreatedEvent)
	return fmt.Sprintf(`
      trigger_on_event {
        type = "PluginPatchsetCreatedEvent"

        exclude_drafts         = %v
        exclude_trivial_rebase = %v
        exclude_no_code_change = %v
        exclude_private_state  = %v
        exclude_wip_state      = %v
      }
`, event.ExcludeDrafts, event.ExcludeTrivialRebase, event.ExcludeNoCodeChange,
		event.ExcludePrivateState, event.ExcludeWipState)
}
