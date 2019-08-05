package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	ensureDefinitionFuncs["*client.CpsScmFlowDefinition"] = ensureCpsScmFlowDefinition
	definitionCodeFuncs["*client.CpsScmFlowDefinition"] = cpsScmFlowDefinitionCode
}

func ensureCpsScmFlowDefinition(definitionInterface client.JobDefinition) error {
	definition := definitionInterface.(*client.CpsScmFlowDefinition)
	if err := ensureGitUserRemoteConfigs(definition.SCM.UserRemoteConfigs); err != nil {
		return err
	}
	if err := ensureGitScmBranches(definition.SCM.Branches); err != nil {
		return err
	}
	if err := ensureGitScmSubmoduleConfig(definition.SCM.SubmoduleCfg); err != nil {
		return err
	}
	return ensureGitScmExtensions(definition.SCM.Extensions)
}

func cpsScmFlowDefinitionCode(definitionInterface client.JobDefinition) string {
	definition := definitionInterface.(*client.CpsScmFlowDefinition)
	return fmt.Sprintf(`
resource "jenkins_job_git_scm" "main" {
	job = "${jenkins_job.main.id}"

	config_version = "%v"
	script_path    = "%v"
	lightweight    = %v
}
`, definition.SCM.ConfigVersion, definition.ScriptPath, definition.Lightweight) +
		jobGitScmUserRemoteConfigsCode(definition.SCM.UserRemoteConfigs) +
		jobGitScmBranchesCode(definition.SCM.Branches) +
		jobGitScmExtensionsCode(definition.SCM.Extensions)
}
