package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmExtension interface {
	fromClientExtension(client.GitScmExtension) (jobGitScmExtension, error)
	toClientExtension() (client.GitScmExtension, error)
}
