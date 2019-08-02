package client

import "encoding/xml"

func init() {
	jobBuildDiscarderPropertyStrategyUnmarshalFunc["hudson.tasks.LogRotator"] = unmarshalJobBuildDiscarderPropertyStrategyLogRotator
}

type JobBuildDiscarderPropertyStrategyLogRotator struct {
	Id    string `xml:"id,attr,omitempty"`
	Class string `xml:"class,attr"`

	DaysToKeep         int `xml:"daysToKeep"`
	NumToKeep          int `xml:"numToKeep"`
	ArtifactDaysToKeep int `xml:"artifactDaysToKeep"`
	ArtifactNumToKeep  int `xml:"artifactNumToKeep"`
}

func NewJobBuildDiscarderPropertyStrategyLogRotator() *JobBuildDiscarderPropertyStrategyLogRotator {
	return &JobBuildDiscarderPropertyStrategyLogRotator{
		Class: "hudson.tasks.LogRotator",

		DaysToKeep:         -1,
		NumToKeep:          -1,
		ArtifactDaysToKeep: -1,
		ArtifactNumToKeep:  -1,
	}
}

func (s *JobBuildDiscarderPropertyStrategyLogRotator) GetId() string {
	return s.Id
}

func (s *JobBuildDiscarderPropertyStrategyLogRotator) SetId(id string) {
	s.Id = id
}

func unmarshalJobBuildDiscarderPropertyStrategyLogRotator(d *xml.Decoder, start xml.StartElement) (JobBuildDiscarderPropertyStrategy, error) {
	strategy := NewJobBuildDiscarderPropertyStrategyLogRotator()
	err := d.DecodeElement(strategy, &start)
	if err != nil {
		return nil, err
	}
	return strategy, nil
}
