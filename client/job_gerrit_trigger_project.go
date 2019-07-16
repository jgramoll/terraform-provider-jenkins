package client

type JobGerritTriggerProjects struct {
	Items *[]*JobGerritTriggerProject `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.data.GerritProject"`
}

type JobGerritTriggerProject struct {
	CompareType string                    `xml:"compareType"`
	Pattern     string                    `xml:"pattern"`
	Branches    *JobGerritTriggerBranches `xml:"branches"`

	DisableStrictForbiddenFileVerification bool `xml:"disableStrictForbiddenFileVerification"`
}
