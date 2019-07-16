package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritBranch struct {
	Project string `mapstructure:"project"`
	// TODO enum / validation
	CompareType string           `mapstructure:"compare_type"`
	Pattern     string           `mapstructure:"pattern"`
}

func newJobGerritBranch() *jobGerritProject {
	return &jobGerritProject{}
}

func (branch *jobGerritBranch) toClientProperty() *client.GitBranchSpec {
	return &client.GitBranchSpec{
		// TODO
		// Id: branch.RefId,
	}
}

func (config *jobGerritBranch) setResourceData(d *schema.ResourceData) error {
	// if err := d.Set("name", config.Name); err != nil {
	// 	return err
	// }
	return nil
	// return d.Set("credentials_id", config.CredentialsId)
}
