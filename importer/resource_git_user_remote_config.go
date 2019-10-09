package main

import (
	"fmt"
	"strings"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func jobGitScmUserRemoteConfigsCode(configs *client.GitUserRemoteConfigs) string {
	code := ""
	for _, item := range *configs.Items {
		// Need $ -> $$ for tf escape
		refspec := strings.ReplaceAll(item.Refspec, "$", "$$")
		code += fmt.Sprintf(`
      user_remote_config {
        refspec        = "%s"
        url            = "%s"
        credentials_id = "%s"
      }
`, refspec, item.Url, item.CredentialsId)
	}
	return code
}
