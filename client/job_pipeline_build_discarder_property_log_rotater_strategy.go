package client

type JobPipelineBuildDiscarderPropertyStrategyLogRotator struct {
	Class string `xml:"class,attr"`

	DaysToKeep         int `xml:"daysToKeep"`
	NumToKeep          int `xml:"numToKeep"`
	ArtifactDaysToKeep int `xml:"artifactDaysToKeep"`
	ArtifactNumToKeep  int `xml:"artifactNumToKeep"`
}

func NewJobPipelineBuildDiscarderPropertyStrategyLogRotator() *JobPipelineBuildDiscarderPropertyStrategyLogRotator {
	return &JobPipelineBuildDiscarderPropertyStrategyLogRotator{
		Class: "hudson.tasks.LogRotator",

		DaysToKeep:         -1,
		NumToKeep:          -1,
		ArtifactDaysToKeep: -1,
		ArtifactNumToKeep:  -1,
	}
}
