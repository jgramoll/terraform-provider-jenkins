package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobGitScmExtensionCodeFuncs["*client.GitScmCleanBeforeCheckoutExtension"] = jobGitScmCleanBeforeCheckoutExtensionCode
}

func jobGitScmCleanBeforeCheckoutExtensionCode(client.GitScmExtension) string {
	return `
resource "jenkins_job_git_scm_clean_before_checkout_extension" "main" {
	scm = "${jenkins_job_git_scm.main.id}"
}
`
}
