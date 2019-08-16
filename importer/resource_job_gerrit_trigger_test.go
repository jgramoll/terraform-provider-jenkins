package main

import (
	"fmt"
	"testing"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"

	// "github.com/jgramoll/terraform-provider-jenkins/provider"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func testJobPipelineTriggersProperties() *[]client.JobProperty {
	gerritBranch := client.NewJobGerritTriggerBranch()
	gerritBranch.CompareType = client.CompareTypeRegExp
	gerritBranch.Pattern = "my-branch"
	gerritFilePath := client.NewJobGerritTriggerFilePath()
	gerritFilePath.CompareType = client.CompareTypeRegExp
	gerritFilePath.Pattern = "my-file-path"
	gerritProject := client.NewJobGerritTriggerProject()
	gerritProject.CompareType = client.CompareTypePlain
	gerritProject.Pattern = "my-project"
	gerritProject.Branches = gerritProject.Branches.Append(gerritBranch)
	gerritProject.FilePaths = gerritProject.FilePaths.Append(gerritFilePath)
	gerritTrigger := client.NewJobGerritTrigger()
	gerritTrigger.Plugin = "gerrit-trigger@2.29.0"
	gerritTrigger.Projects = gerritTrigger.Projects.Append(gerritProject)
	gerritTriggerChangeMergedEvent := client.NewJobGerritTriggerPluginChangeMergedEvent()
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerChangeMergedEvent)
	gerritTriggerPatchsetEvent := client.NewJobGerritTriggerPluginPatchsetCreatedEvent()
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerPatchsetEvent)
	gerritTriggerDraftEvent := client.NewJobGerritTriggerPluginDraftPublishedEvent()
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerDraftEvent)

	triggerJobProperty := client.NewJobPipelineTriggersProperty()
	triggerJobProperty.Triggers = triggerJobProperty.Triggers.Append(gerritTrigger)

	triggerJobProperty2 := client.NewJobPipelineTriggersProperty()
	gerritTrigger2 := client.NewJobGerritTrigger()
	triggerJobProperty2.Triggers = triggerJobProperty2.Triggers.Append(gerritTrigger2)
	return &[]client.JobProperty{triggerJobProperty, triggerJobProperty2}
}

func TestEnsureJobGerritTrigger(t *testing.T) {
	job := client.NewJob()
	job.Properties.Items = testJobPipelineTriggersProperties()

	if err := ensureJob(job); err != nil {
		t.Fatal(err)
	}
	// TODO
}

func TestJobGerritTriggerCode(t *testing.T) {
	job := client.NewJob()
	job.Properties.Items = testJobPipelineTriggersProperties()

	result := jobCode(job)
	expected := `resource "jenkins_job" "main" {
	name     = ""
	plugin   = ""
	disabled = false
}

resource "jenkins_job_pipeline_triggers_property" "property_1" {
	job = "${jenkins_job.main.name}"
}

resource "jenkins_job_gerrit_trigger" "trigger_1_1" {
	property = "${jenkins_job_pipeline_triggers_property.property_1.id}"

	plugin            = "gerrit-trigger@2.29.0"
	server_name       = "__ANY__"
	silent_mode       = false
	silent_start_mode = false
	escape_quotes     = true

	name_and_email_parameter_mode = "PLAIN"
	commit_message_parameter_mode = "BASE64"
	change_subject_parameter_mode = "PLAIN"
	comment_text_parameter_mode   = "BASE64"
	dynamic_trigger_configuration = false

	skip_vote = {
		on_successful = false
		on_failed     = false
		on_unstable   = false
		on_not_built  = false
	}
}

resource "jenkins_job_gerrit_trigger_change_merged_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_1_1.id}"
}

resource "jenkins_job_gerrit_trigger_patchset_created_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_1_1.id}"

	exclude_drafts         = false
	exclude_trivial_rebase = false
	exclude_no_code_change = false
	exclude_private_state  = false
	exclude_wip_state      = false
}

resource "jenkins_job_gerrit_trigger_draft_published_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_1_1.id}"
}

resource "jenkins_job_gerrit_project" "project_1_1_1" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_1_1.id}"

	compare_type = "PLAIN"
	pattern      = "my-project"
}

resource "jenkins_job_gerrit_branch" "branch_1_1_1_1" {
	project = "${jenkins_job_gerrit_project.project_1_1_1.id}"

	compare_type = "REG_EXP"
	pattern      = "my-branch"
}

resource "jenkins_job_gerrit_file_path" "file_path_1_1_1_1" {
	project = "${jenkins_job_gerrit_project.project_1_1_1.id}"

	compare_type = "REG_EXP"
	pattern      = "my-file-path"
}

resource "jenkins_job_pipeline_triggers_property" "property_2" {
	job = "${jenkins_job.main.name}"
}

resource "jenkins_job_gerrit_trigger" "trigger_2_1" {
	property = "${jenkins_job_pipeline_triggers_property.property_2.id}"

	plugin            = ""
	server_name       = "__ANY__"
	silent_mode       = false
	silent_start_mode = false
	escape_quotes     = true

	name_and_email_parameter_mode = "PLAIN"
	commit_message_parameter_mode = "BASE64"
	change_subject_parameter_mode = "PLAIN"
	comment_text_parameter_mode   = "BASE64"
	dynamic_trigger_configuration = false

	skip_vote = {
		on_successful = false
		on_failed     = false
		on_unstable   = false
		on_not_built  = false
	}
}
`
	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		t.Fatalf("job terraform code not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}

func TestJobGerritTriggerImportScript(t *testing.T) {
	job := client.NewJob()
	job.Id = "id"
	job.Name = "name"
	properties := testJobPipelineTriggersProperties()
	job.Properties.Items = properties

	property := (*properties)[0].(*client.JobPipelineTriggersProperty)
	property.Id = "triggerPropertyId"
	triggers1 := *property.Triggers.Items
	gerritTrigger1 := triggers1[0].(*client.JobGerritTrigger)
	events1 := *gerritTrigger1.TriggerOnEvents.Items
	projects1 := *gerritTrigger1.Projects.Items
	branches1 := *projects1[0].Branches.Items
	filePaths1 := *projects1[0].FilePaths.Items

	property2 := (*properties)[1].(*client.JobPipelineTriggersProperty)
	triggers2 := *property2.Triggers.Items
	gerritTrigger2 := triggers2[0].(*client.JobGerritTrigger)

	result := jobImportScript(job)
	expected := fmt.Sprintf(`terraform init

terraform import jenkins_job.main "%v"

terraform import jenkins_job_pipeline_triggers_property.property_1 "%v"

terraform import jenkins_job_gerrit_trigger.trigger_1_1 "%v"

terraform import jenkins_job_gerrit_trigger_change_merged_event.main "%v"

terraform import jenkins_job_gerrit_trigger_patchset_created_event.main "%v"

terraform import jenkins_job_gerrit_trigger_draft_published_event.main "%v"

terraform import jenkins_job_gerrit_project.project_1_1_1 "%v"

terraform import jenkins_job_gerrit_branch.branch_1_1_1_1 "%v"

terraform import jenkins_job_gerrit_file_path.file_path_1_1_1_1 "%v"

terraform import jenkins_job_pipeline_triggers_property.property_2 "%v"

terraform import jenkins_job_gerrit_trigger.trigger_2_1 "%v"
`, job.Name,
		provider.ResourceJobPropertyId(job.Name, property.GetId()),
		provider.ResourceJobTriggerId(job.Name, property.GetId(), triggers1[0].GetId()),
		provider.ResourceJobGerritTriggerEventId(job.Name, property.GetId(), triggers1[0].GetId(), events1[0].GetId()),
		provider.ResourceJobGerritTriggerEventId(job.Name, property.GetId(), triggers1[0].GetId(), events1[1].GetId()),
		provider.ResourceJobGerritTriggerEventId(job.Name, property.GetId(), triggers1[0].GetId(), events1[2].GetId()),
		provider.ResourceJobGerritProjectId(job.Name, property.GetId(), triggers1[0].GetId(), projects1[0].Id),
		provider.ResourceJobGerritBranchId(job.Name, property.GetId(), triggers1[0].GetId(), projects1[0].Id, branches1[0].Id),
		provider.ResourceJobGerritFilePathId(job.Name, property.GetId(), triggers1[0].GetId(), projects1[0].Id, filePaths1[0].Id),
		provider.ResourceJobPropertyId(job.Name, property2.GetId()),
		provider.ResourceJobTriggerId(job.Name, property2.GetId(), gerritTrigger2.Id))

	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		t.Fatalf("job terraform import script not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}
