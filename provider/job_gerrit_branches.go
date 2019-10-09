package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritBranches []*jobGerritBranch

func (branches *jobGerritBranches) toClientBranches() (*client.JobGerritTriggerBranches, error) {
	clientBranches := client.NewJobGerritTriggerBranches()
	for _, branch := range *branches {
		clientBranch, err := branch.toClientBranch()
		if err != nil {
			return nil, err
		}
		clientBranches = clientBranches.Append(clientBranch)
	}
	return clientBranches, nil
}

func (*jobGerritBranches) fromClientBranches(clientBranches *client.JobGerritTriggerBranches) *jobGerritBranches {
	if clientBranches == nil || clientBranches.Items == nil {
		return nil
	}
	branches := jobGerritBranches{}
	for _, clientBranch := range *clientBranches.Items {
		branch := newJobGerritBranchFromClient(clientBranch)
		branches = append(branches, branch)
	}
	return &branches
}
