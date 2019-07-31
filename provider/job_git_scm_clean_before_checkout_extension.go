package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmCleanBeforeCheckoutExtension struct {
	Scm string `mapstructure:"scm"`
}

func newJobGitScmCleanBeforeCheckoutExtension() *jobGitScmCleanBeforeCheckoutExtension {
	return &jobGitScmCleanBeforeCheckoutExtension{}
}

func (branch *jobGitScmCleanBeforeCheckoutExtension) fromClientExtension(clientExtension client.GitScmExtension) jobGitScmExtension {
	return newJobGitScmCleanBeforeCheckoutExtension()
}

func (branch *jobGitScmCleanBeforeCheckoutExtension) toClientExtension(extensionId string) (client.GitScmExtension, error) {
	clientExtension := client.NewGitScmCleanBeforeCheckoutExtension()
	clientExtension.Id = extensionId
	return clientExtension, nil
}

func (config *jobGitScmCleanBeforeCheckoutExtension) setResourceData(d *schema.ResourceData) error {
	return nil
}
