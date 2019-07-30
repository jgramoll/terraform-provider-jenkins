package client

import (
	"encoding/xml"
)

type GitScmBranches struct {
	XMLName xml.Name             `xml:"branches"`
	Items   *[]*GitScmBranchSpec `xml:"hudson.plugins.git.BranchSpec"`
}

func NewGitScmBranches() *GitScmBranches {
	return &GitScmBranches{
		Items: &[]*GitScmBranchSpec{},
	}
}

func (branches *GitScmBranches) Append(branch *GitScmBranchSpec) *GitScmBranches {
	newBranches := NewGitScmBranches()
	if branches.Items != nil {
		*newBranches.Items = append(*branches.Items, branch)
	} else {
		*newBranches.Items = []*GitScmBranchSpec{branch}
	}
	return newBranches
}
