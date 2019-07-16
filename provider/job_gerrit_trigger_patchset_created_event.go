package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerPatchSetCreatedEvent struct {
	Job                  string `mapstructure:"job"`
	Trigger              string `mapstructure:"trigger"`
	ExcludeDrafts        bool   `mapstructure:"exclude_drafts"`
	ExcludeTrivialRebase bool   `mapstructure:"exclude_trivial_rebase"`
	ExcludeNoCodeChange  bool   `mapstructure:"exclude_no_code_change"`
	ExcludePrivateState  bool   `mapstructure:"exclude_private_state"`
	ExcludeWipState      bool   `mapstructure:"exclude_wip_state"`
}

func newJobGerritTriggerPatchSetCreatedEvent() *jobGerritTriggerPatchSetCreatedEvent {
	return &jobGerritTriggerPatchSetCreatedEvent{}
}

func (branch *jobGerritTriggerPatchSetCreatedEvent) toClientEvent() *client.GitBranchSpec {
	return &client.GitBranchSpec{
		// TODO
		// Id: branch.RefId,
	}
}

func (config *jobGerritTriggerPatchSetCreatedEvent) setResourceData(d *schema.ResourceData) error {
	// if err := d.Set("name", config.Name); err != nil {
	// 	return err
	// }
	return nil
	// return d.Set("credentials_id", config.CredentialsId)
}
