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
	var newBranchItems []*GitScmBranchSpec
	if branches.Items != nil {
		newBranchItems = append(*branches.Items, branch)
	} else {
		newBranchItems = append(newBranchItems, branch)
	}
	newBranches := NewGitScmBranches()
	newBranches.Items = &newBranchItems
	return newBranches
}
