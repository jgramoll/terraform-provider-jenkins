package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmBranches []*jobGitScmBranch

func (branches *jobGitScmBranches) toClientBranches() *client.GitScmBranches {
	clientBranches := client.NewGitScmBranches()
	for _, branch := range *branches {
		clientBranch := client.NewGitScmBranchSpec()
		clientBranch.Name = branch.Name
		clientBranches = clientBranches.Append(clientBranch)
	}
	return clientBranches
}

func (*jobGitScmBranches) fromClientBranches(clientBranches *client.GitScmBranches) *jobGitScmBranches {
	if clientBranches == nil || clientBranches.Items == nil {
		return nil
	}
	branches := jobGitScmBranches{}
	for _, clientBranch := range *clientBranches.Items {
		branch := newJobGitScmBranch()
		branch.Name = clientBranch.Name
		branches = append(branches, branch)
	}
	return &branches
}
