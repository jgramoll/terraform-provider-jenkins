package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	ensureDefinitionFuncs["*client.CpsScmFlowDefinition"] = ensureCpsScmFlowDefinition
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
