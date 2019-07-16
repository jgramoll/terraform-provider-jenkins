package client

type JobGerritTriggerBranches struct {
	Items *[]*JobGerritTriggerBranch `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.data.Branch"`
}

type JobGerritTriggerBranch struct {
	CompareType string `xml:"compareType"`
	Pattern     string `xml:"pattern"`
}
