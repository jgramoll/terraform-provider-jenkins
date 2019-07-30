package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritBranch struct {
	CompareType string `mapstructure:"compare_type"`
	Pattern     string `mapstructure:"pattern"`
}

func newJobGerritBranch() *jobGerritBranch {
	return &jobGerritBranch{}
}

func newJobGerritBranchFromClient(clientBranch *client.JobGerritTriggerBranch) *jobGerritBranch {
	branch := newJobGerritBranch()
	branch.CompareType = clientBranch.CompareType.String()
	branch.Pattern = clientBranch.Pattern
	return branch
}

func (branch *jobGerritBranch) toClientBranch(branchId string) (*client.JobGerritTriggerBranch, error) {
	clientBranch := client.NewJobGerritTriggerBranch()
	clientBranch.Id = branchId
	compareType, err := client.ParseCompareType(branch.CompareType)
	if err != nil {
		return nil, err
	}
	clientBranch.CompareType = compareType
	clientBranch.Pattern = branch.Pattern
	return clientBranch, nil
}

func (branch *jobGerritBranch) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("compare_type", branch.CompareType); err != nil {
		return err
	}
	return d.Set("pattern", branch.Pattern)
}
