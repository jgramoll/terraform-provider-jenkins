package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmCleanBeforeCheckoutExtension struct {
	Job string `mapstructure:"job"`
	Scm string `mapstructure:"scm"`
}

func newJobGitScmCleanBeforeCheckoutExtension() *jobGitScmCleanBeforeCheckoutExtension {
	return &jobGitScmCleanBeforeCheckoutExtension{}
}

func (branch *jobGitScmCleanBeforeCheckoutExtension) toClientExtension() *client.GitBranchSpec {
	return &client.GitBranchSpec{
		// Id: branch.RefId,
	}
}

func (config *jobGitScmCleanBeforeCheckoutExtension) setResourceData(d *schema.ResourceData) error {
	// if err := d.Set("name", config.Name); err != nil {
	// 	return err
	// }
	return nil
	// return d.Set("credentials_id", config.CredentialsId)
}
