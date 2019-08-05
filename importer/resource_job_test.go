package main

import (
	"testing"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestJobCode(t *testing.T) {
	job := client.NewJob()
	job.Name = "Premerge checks"
	job.Plugin = "flow-plugin"
	job.Description = "my-desc"
	job.Actions = job.Actions.Append(client.NewJobDeclarativeJobAction())
	job.Actions = job.Actions.Append(client.NewJobDeclarativeJobPropertyTrackerAction())

	definition := client.NewCpsScmFlowDefinition()
	definition.ScriptPath = "my-Jenkinsfile"
	definition.SCM = client.NewGitScm()
	definition.SCM.ConfigVersion = "my-version"

	remoteConfig := client.NewGitUserRemoteConfig()
	remoteConfig.Refspec = "refspec"
	remoteConfig.Url = "url.to.server"
	remoteConfig.CredentialsId = "creds"
	definition.SCM.UserRemoteConfigs = definition.SCM.UserRemoteConfigs.Append(remoteConfig)

	scmExtension := client.NewGitScmCleanBeforeCheckoutExtension()
	scmExtension.Id = "extension-id"
	definition.SCM.Extensions = definition.SCM.Extensions.Append(scmExtension)

	branchSpec := client.NewGitScmBranchSpec()
	branchSpec.Name = "branchspec"
	definition.Id = "definition-id"
	definition.SCM.Branches = definition.SCM.Branches.Append(branchSpec)
	job.Definition = definition

	gerritBranch := client.NewJobGerritTriggerBranch()
	gerritBranch.CompareType = client.CompareTypeRegExp
	gerritBranch.Pattern = "my-branch"
	gerritProject := client.NewJobGerritTriggerProject()
	gerritProject.CompareType = client.CompareTypePlain
	gerritProject.Pattern = "my-project"
	gerritProject.Branches = gerritProject.Branches.Append(gerritBranch)
	gerritTrigger := client.NewJobGerritTrigger()
	gerritTrigger.Plugin = "gerrit-trigger@2.29.0"
	gerritTrigger.Projects = gerritTrigger.Projects.Append(gerritProject)
	gerritTriggerPatchsetEvent := client.NewJobGerritTriggerPluginPatchsetCreatedEvent()
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerPatchsetEvent)
	gerritTriggerDraftEvent := client.NewJobGerritTriggerPluginDraftPublishedEvent()
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerDraftEvent)
	triggerJobProperty := client.NewJobPipelineTriggersProperty()
	triggerJobProperty.Id = "trigger-id"
	triggerJobProperty.Triggers = triggerJobProperty.Triggers.Append(gerritTrigger)
	job.Properties = job.Properties.Append(triggerJobProperty)

	discardPropertyStrategy := client.NewJobBuildDiscarderPropertyStrategyLogRotator()
	discardPropertyStrategy.DaysToKeep = 1
	discardPropertyStrategy.NumToKeep = 2
	discardPropertyStrategy.ArtifactDaysToKeep = 3
	discardPropertyStrategy.ArtifactNumToKeep = 4
	discardProperty := client.NewJobBuildDiscarderProperty()
	discardProperty.Id = "discard-id"
	discardProperty.Strategy.Item = discardPropertyStrategy
	job.Properties = job.Properties.Append(discardProperty)

	datadogJobProperty := client.NewJobDatadogJobProperty()
	datadogJobProperty.Plugin = "datadog@0.7.1"
	job.Properties = job.Properties.Append(datadogJobProperty)

	jiraProjectProperty := client.NewJobJiraProjectProperty()
	job.Properties = job.Properties.Append(jiraProjectProperty)

	result := jobCode(job)
	expected := `resource "jenkins_job" "main" {
	name     = "Premerge checks"
	plugin   = "flow-plugin"
	disabled = false
}

resource "jenkins_job_declarative_job_action" "main" {
	job = "${jenkins_job.main.id}"
}

resource "jenkins_job_declarative_job_property_tracker_action" "main" {
	job = "${jenkins_job.main.id}"
}

resource "jenkins_job_git_scm" "main" {
	job = "${jenkins_job.main.id}"

	config_version = "my-version"
	script_path    = "my-Jenkinsfile"
	lightweight    = false
}

resource "jenkins_job_git_scm_user_remote_config" "config_1" {
	scm = "${jenkins_job_git_scm.main.id}"

	refspec        = "refspec"
	url            = "url.to.server"
	credentials_id = "creds"
}

resource "jenkins_job_git_scm_branch" "branch_1" {
	scm = "${jenkins_job_git_scm.main.id}"

	name = "branchspec"
}

resource "jenkins_job_git_scm_clean_before_checkout_extension" "main" {
	scm = "${jenkins_job_git_scm.main.id}"
}

resource "jenkins_job_pipeline_triggers_property" "main" {
	job = "${jenkins_job.main.id}"
}

resource "jenkins_job_gerrit_trigger" "main" {
	property = "${jenkins_job_pipeline_triggers_property.main.id}"

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

resource "jenkins_job_gerrit_trigger_patchset_created_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.main.id}"

	exclude_drafts         = false
	exclude_trivial_rebase = false
	exclude_no_code_change = false
	exclude_private_state  = false
	exclude_wip_state      = false
}

resource "jenkins_job_gerrit_trigger_draft_published_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.main.id}"
}

resource "jenkins_job_gerrit_project" "project_1" {
	trigger = "${jenkins_job_gerrit_trigger.main.id}"

	compare_type = "PLAIN"
	pattern      = "my-project"
}

resource "jenkins_job_gerrit_branch" "branch_1" {
	project = "${jenkins_job_gerrit_project.main.id}"

	compare_type = "REG_EXP"
	pattern      = "my-branch"
}

resource "jenkins_job_build_discarder_property" "main" {
	job = "${jenkins_job.main.id}"
}

resource "jenkins_job_build_discarder_property_log_rotator_strategy" "main" {
	property = "${jenkins_job_build_discarder_property.main.id}"

	days_to_keep          = "1"
	num_to_keep           = "2"
	artifact_days_to_keep = "3"
	artifact_num_to_keep  = "4"
}

resource "jenkins_job_datadog_job_property" "main" {
	job = "${jenkins_job.main.id}"
}

resource "jenkins_job_jira_project_property" "main" {
	job = "${jenkins_job.main.id}"
}
`
	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		for _, d := range diffs {
			println("DIFF", d.Text)
		}
		t.Fatalf("job definition not expected: %s", "SDF")
	}
}
