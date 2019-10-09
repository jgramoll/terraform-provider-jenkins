package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmBranch struct {
	Name string `mapstructure:"name"`
}

func newJobGitScmBranch() *jobGitScmBranch {
	return &jobGitScmBranch{}
}

func newGitScmBranchFromClient(clientBranch *client.GitScmBranchSpec) *jobGitScmBranch {
	branch := newJobGitScmBranch()
	branch.Name = clientBranch.Name
	return branch
}

func (branch *jobGitScmBranch) toClientBranch() *client.GitScmBranchSpec {
	clientBranch := client.NewGitScmBranchSpec()
	clientBranch.Name = branch.Name
	return clientBranch
}
