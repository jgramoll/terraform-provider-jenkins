package provider

import (
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

func (branch *jobGerritBranch) toClientBranch() (*client.JobGerritTriggerBranch, error) {
	clientBranch := client.NewJobGerritTriggerBranch()
	compareType, err := client.ParseCompareType(branch.CompareType)
	if err != nil {
		return nil, err
	}
	clientBranch.CompareType = compareType
	clientBranch.Pattern = branch.Pattern
	return clientBranch, nil
}
