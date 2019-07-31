package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
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

func (branch *jobGitScmBranch) toClientBranch(branchId string) *client.GitScmBranchSpec {
	clientBranch := client.NewGitScmBranchSpec()
	clientBranch.Id = branchId
	clientBranch.Name = branch.Name
	return clientBranch
}

func (branch *jobGitScmBranch) setResourceData(d *schema.ResourceData) error {
	return d.Set("name", branch.Name)
}
