package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobGerritTriggerOnEventsCodeFuncs["*client.JobGerritTriggerPluginChangeMergedEvent"] = jobGerritTriggerPluginChangeMergedEventCode
}

func jobGerritTriggerPluginChangeMergedEventCode(
	triggerIndex string, e client.JobGerritTriggerOnEvent,
) string {
	return fmt.Sprintf(`
resource "jenkins_job_gerrit_trigger_change_merged_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_%v.id}"
}
`, triggerIndex)
}
