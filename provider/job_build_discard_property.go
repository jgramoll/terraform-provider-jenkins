package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobBuildDiscardProperty struct {
	Job string `mapstructure:"job"`
	Scm string `mapstructure:"scm"`
}

func newJobBuildDiscardProperty() *jobBuildDiscardProperty {
	return &jobBuildDiscardProperty{}
}

func (branch *jobBuildDiscardProperty) toClientProperty() *client.GitBranchSpec {
	return &client.GitBranchSpec{
		// TODO
		// Id: branch.RefId,
	}
}

func (config *jobBuildDiscardProperty) setResourceData(d *schema.ResourceData) error {
	// if err := d.Set("name", config.Name); err != nil {
	// 	return err
	// }
	return nil
	// return d.Set("credentials_id", config.CredentialsId)
}
