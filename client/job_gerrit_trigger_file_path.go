package client

type JobGerritTriggerFilePath struct {
	Id          string      `xml:"id,attr,omitempty"`
	CompareType CompareType `xml:"compareType"`
	Pattern     string      `xml:"pattern"`
}

func NewJobGerritTriggerFilePath() *JobGerritTriggerFilePath {
	return &JobGerritTriggerFilePath{}
}
