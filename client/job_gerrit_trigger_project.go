package client

import (
	"encoding/xml"
	"errors"
)

// ErrJobGerritTriggerBranchNotFound job trigger gerrit branch not found
var ErrJobGerritTriggerBranchNotFound = errors.New("Could not find job trigger gerrit branch")

type JobGerritTriggerProject struct {
	XMLName xml.Name `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.data.GerritProject"`
	Id      string   `xml:"id,attr,omitempty"`

	CompareType CompareType               `xml:"compareType"`
	Pattern     string                    `xml:"pattern"`
	Branches    *JobGerritTriggerBranches `xml:"branches"`

	DisableStrictForbiddenFileVerification bool `xml:"disableStrictForbiddenFileVerification"`
}

func NewJobGerritTriggerProject() *JobGerritTriggerProject {
	return &JobGerritTriggerProject{
		Branches: NewJobGerritTriggerBranches(),
	}
}

func (project *JobGerritTriggerProject) GetBranch(branchId string) (*JobGerritTriggerBranch, error) {
	branches := *(project.Branches).Items
	for _, branch := range branches {
		if branch.Id == branchId {
			return branch, nil
		}
	}
	return nil, ErrJobGerritTriggerBranchNotFound
}

func (project *JobGerritTriggerProject) UpdateBranch(branch *JobGerritTriggerBranch) error {
	oldBranch, err := project.GetBranch(branch.Id)
	if err != nil {
		return err
	}
	*oldBranch = *branch
	return nil
}

func (project *JobGerritTriggerProject) DeleteBranch(branchId string) error {
	branches := *project.Branches.Items
	for i, branch := range branches {
		if branch.Id == branchId {
			*project.Branches.Items = append(branches[:i], branches[i+1:]...)
			return nil
		}
	}
	return ErrJobGerritTriggerBranchNotFound
}
