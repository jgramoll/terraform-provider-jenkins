package client

import (
	"encoding/xml"
)

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
