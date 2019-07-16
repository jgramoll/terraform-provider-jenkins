package client

type JobGerritTriggerBranch struct {
	CompareType CompareType `xml:"compareType"`
	Pattern     string `xml:"pattern"`
}

func NewJobGerritTriggerBranch() *JobGerritTriggerBranch {
	return &JobGerritTriggerBranch{}
}
