package main

import (
	"testing"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func testNewJob() *client.Job {
	job := client.NewJob()
	job.Name = "Premerge checks"
	job.Plugin = "flow-plugin"
	job.Description = "my-desc"
	declarativeJobAction := client.NewJobDeclarativeJobAction()
	declarativeJobAction.Plugin = "DeclarativeJobPlugin"
	job.Actions = job.Actions.Append(declarativeJobAction)
	declarativeJobPropertyTrackerAction := client.NewJobDeclarativeJobPropertyTrackerAction()
	declarativeJobPropertyTrackerAction.Plugin = "DeclarativeJobTrackerPlugin"
	job.Actions = job.Actions.Append(declarativeJobPropertyTrackerAction)

	definition := client.NewCpsScmFlowDefinition()
	definition.Plugin = "gitPlugin"
	definition.ScriptPath = "my-Jenkinsfile"
	definition.SCM = client.NewGitScm()
	definition.SCM.Plugin = "gitScmPlugin"
	definition.SCM.ConfigVersion = "my-version"

	remoteConfig := client.NewGitUserRemoteConfig()
	remoteConfig.Refspec = "${GERRIT_REFSPEC}"
	remoteConfig.Url = "url.to.server"
	remoteConfig.CredentialsId = "creds"
	definition.SCM.UserRemoteConfigs = definition.SCM.UserRemoteConfigs.Append(remoteConfig)

	scmExtension := client.NewGitScmCleanBeforeCheckoutExtension()
	definition.SCM.Extensions = definition.SCM.Extensions.Append(scmExtension)

	branchSpec := client.NewGitScmBranchSpec()
	branchSpec.Name = "branchspec"
	definition.SCM.Branches = definition.SCM.Branches.Append(branchSpec)
	job.Definition = definition

	discardPropertyStrategy := client.NewJobBuildDiscarderPropertyStrategyLogRotator()
	discardPropertyStrategy.DaysToKeep = 1
	discardPropertyStrategy.NumToKeep = 2
	discardPropertyStrategy.ArtifactDaysToKeep = 3
	discardPropertyStrategy.ArtifactNumToKeep = 4
	discardProperty := client.NewJobBuildDiscarderProperty()
	discardProperty.Strategy.Item = discardPropertyStrategy
	job.Properties = job.Properties.Append(discardProperty)

	datadogJobProperty := client.NewJobDatadogJobProperty()
	datadogJobProperty.Plugin = "datadog@0.7.1"
	job.Properties = job.Properties.Append(datadogJobProperty)

	jiraProjectProperty := client.NewJobJiraProjectProperty()
	jiraProjectProperty.Plugin = "jiraPlugin"
	job.Properties = job.Properties.Append(jiraProjectProperty)
	return job
}

func TestJobCode(t *testing.T) {
	job := testNewJob()

	result := jobCode(job)
	expected := `resource "jenkins_job" "main" {
  name     = "Premerge checks"
  plugin   = "flow-plugin"
  disabled = false

  action {
    type = "DeclarativeJobAction"
    plugin = "DeclarativeJobPlugin"
  }

  action {
    type = "DeclarativeJobPropertyTrackerAction"
    plugin = "DeclarativeJobTrackerPlugin"
  }

  definition {
    type   = "CpsScmFlowDefinition"
    plugin = "gitPlugin"

    scm {
      type = "GitSCM"
      plugin = "gitScmPlugin"

      config_version = "my-version"
      script_path    = "my-Jenkinsfile"
      lightweight    = false

      user_remote_config {
        refspec        = "$${GERRIT_REFSPEC}"
        url            = "url.to.server"
        credentials_id = "creds"
      }

      branch {
        name = "branchspec"
      }

      extension {
        type = "CleanBeforeCheckout"
      }
    }
  }

  property {
    type = "BuildDiscarderProperty"

    strategy {
      type  = "LogRotator"

      days_to_keep          = "1"
      num_to_keep           = "2"
      artifact_days_to_keep = "3"
      artifact_num_to_keep  = "4"
    }
  }

  property {
    type = "DatadogJobProperty"
    plugin="datadog@0.7.1"

    emit_on_checkout = false
  }

  property {
    type = "JiraProjectProperty"
    plugin="jiraPlugin"
  }
}
`
	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		t.Fatalf("job terraform code not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}

func TestJobImportScript(t *testing.T) {
	job := testNewJob()

	result := jobImportScript(job)
	expected := `
terraform init

terraform import jenkins_job.main "Premerge checks"
`
	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		t.Fatalf("job terraform import script not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}
