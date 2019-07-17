package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritProject struct {
	Trigger string `mapstructure:"trigger"`
	// TODO enum / validation
	CompareType string `mapstructure:"compare_type"`
	Pattern     string `mapstructure:"pattern"`
}

func newJobGerritProject() *jobGerritProject {
	return &jobGerritProject{}
}

func (branch *jobGerritProject) toClientProperty() *client.GitScmBranchSpec {
	return &client.GitScmBranchSpec{
		// TODO
		// Id: branch.RefId,
	}
}

func (config *jobGerritProject) setResourceData(d *schema.ResourceData) error {
	// if err := d.Set("name", config.Name); err != nil {
	// 	return err
	// }
	return nil
	// return d.Set("credentials_id", config.CredentialsId)
}
