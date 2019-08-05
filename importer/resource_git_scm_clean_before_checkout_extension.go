package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	jobGitScmExtensionCodeFuncs["*client.GitScmCleanBeforeCheckoutExtension"] = jobGitScmCleanBeforeCheckoutExtensionCode
	jobGitScmExtensionImportScriptFuncs["*client.GitScmCleanBeforeCheckoutExtension"] = jobGitScmCleanBeforeCheckoutExtensionImportScript
}

func jobGitScmCleanBeforeCheckoutExtensionCode(client.GitScmExtension) string {
	return `
resource "jenkins_job_git_scm_clean_before_checkout_extension" "main" {
	scm = "${jenkins_job_git_scm.main.id}"
}
`
}

func jobGitScmCleanBeforeCheckoutExtensionImportScript(jobName string, definitionId string, extension client.GitScmExtension) string {
	return fmt.Sprintf(`
terraform import jenkins_job_git_scm_clean_before_checkout_extension.main "%v"
`, provider.ResourceJobGitScmExtensionId(jobName, definitionId, extension.GetId()))
}
