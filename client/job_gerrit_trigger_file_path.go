package client

type JobGerritTriggerFilePath struct {
	CompareType CompareType `xml:"compareType"`
	Pattern     string      `xml:"pattern"`
}

func NewJobGerritTriggerFilePath() *JobGerritTriggerFilePath {
	return &JobGerritTriggerFilePath{}
}
