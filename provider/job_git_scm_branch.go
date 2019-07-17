package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmBranch struct {
	Job  string `mapstructure:"job"`
	Name string `mapstructure:"name"`
}

func newJobGitScmBranch() *jobGitScmBranch {
	return &jobGitScmBranch{}
}

func (branch *jobGitScmBranch) toClientBranch() *client.GitScmBranchSpec {
	return &client.GitScmBranchSpec{
		// Id:   branch.RefId,
		Name: branch.Name,
	}
}

func (config *jobGitScmBranch) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("name", config.Name); err != nil {
		return err
	}
	return nil
	// return d.Set("credentials_id", config.CredentialsId)
}
