package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobGitScmExtensionCodeFuncs["*client.GitScmCleanBeforeCheckoutExtension"] = jobGitScmCleanBeforeCheckoutExtensionCode
}

func jobGitScmCleanBeforeCheckoutExtensionCode(client.GitScmExtension) string {
	return `
      extension {
        type = "CleanBeforeCheckout"
      }
`
}
