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
  definition {
    type   = "CpsScmFlowDefinition"
    plugin = "%s"

    script_path = "%s"
    lightweight = %v

    scm {
      type   = "GitSCM"
      plugin = "%s"

      config_version = "%s"
%s%s%s    }
  }
`, definition.Plugin, definition.ScriptPath, definition.Lightweight, definition.SCM.Plugin, definition.SCM.ConfigVersion,
		jobGitScmUserRemoteConfigsCode(definition.SCM.UserRemoteConfigs),
		jobGitScmBranchesCode(definition.SCM.Branches),
		jobGitScmExtensionsCode(definition.SCM.Extensions))
}
