package client

type JobGerritTriggerBranch struct {
	Id          string      `xml:"id,attr"`
	CompareType CompareType `xml:"compareType"`
	Pattern     string      `xml:"pattern"`
}

func NewJobGerritTriggerBranch() *JobGerritTriggerBranch {
	return &JobGerritTriggerBranch{}
}
