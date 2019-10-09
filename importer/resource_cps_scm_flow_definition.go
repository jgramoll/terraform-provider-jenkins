package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	definitionCodeFuncs["*client.CpsScmFlowDefinition"] = cpsScmFlowDefinitionCode
}

func cpsScmFlowDefinitionCode(definitionInterface client.JobDefinition) string {
	definition := definitionInterface.(*client.CpsScmFlowDefinition)
	return fmt.Sprintf(`
resource "jenkins_job_git_scm" "main" {
	job = "${jenkins_job.main.name}"

	plugin     = "%v"
	git_plugin = "%v"

	config_version = "%v"
	script_path    = "%v"
	lightweight    = %v
}
`, definition.Plugin, definition.SCM.Plugin, definition.SCM.ConfigVersion, definition.ScriptPath, definition.Lightweight) +
		jobGitScmUserRemoteConfigsCode(definition.SCM.UserRemoteConfigs) +
		jobGitScmBranchesCode(definition.SCM.Branches) +
		jobGitScmExtensionsCode(definition.SCM.Extensions)
}
