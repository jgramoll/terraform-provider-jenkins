package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerDraftPublishedEvent struct {
	Job     string `mapstructure:"job"`
	Trigger string `mapstructure:"trigger"`
}

func newJobGerritTriggerDraftPublishedEvent() *jobGerritTriggerDraftPublishedEvent {
	return &jobGerritTriggerDraftPublishedEvent{}
}

func (branch *jobGerritTriggerDraftPublishedEvent) toClientExtension() *client.GitScmBranchSpec {
	return &client.GitScmBranchSpec{
		// TODO
		// Id: branch.RefId,
	}
}

func (config *jobGerritTriggerDraftPublishedEvent) setResourceData(d *schema.ResourceData) error {
	// if err := d.Set("name", config.Name); err != nil {
	// 	return err
	// }
	return nil
	// return d.Set("credentials_id", config.CredentialsId)
}
