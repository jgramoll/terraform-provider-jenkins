package client

import (
	"encoding/xml"
	"errors"
)

// ErrJobGerritTriggerBranchNotFound job trigger gerrit branch not found
var ErrJobGerritTriggerBranchNotFound = errors.New("Could not find job trigger gerrit branch")

// ErrJobGerritTriggerFilePathNotFound job trigger gerrit file path not found
var ErrJobGerritTriggerFilePathNotFound = errors.New("Could not find job trigger gerrit file path")

type JobGerritTriggerProject struct {
	XMLName xml.Name `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.data.GerritProject"`

	CompareType CompareType                `xml:"compareType"`
	Pattern     string                     `xml:"pattern"`
	Branches    *JobGerritTriggerBranches  `xml:"branches"`
	FilePaths   *JobGerritTriggerFilePaths `xml:"filePaths"`

	DisableStrictForbiddenFileVerification bool `xml:"disableStrictForbiddenFileVerification"`
}

func NewJobGerritTriggerProject() *JobGerritTriggerProject {
	return &JobGerritTriggerProject{
		Branches:  NewJobGerritTriggerBranches(),
		FilePaths: NewJobGerritTriggerFilePaths(),
	}
}
