package client

import "encoding/xml"

type JobConfigDeclarativeJobPropertyTrackerAction struct {
	XMLName xml.Name `xml:"org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction"`
	Id      string   `xml:"id,attr,omitempty"`

	JobProperties string `xml:"jobProperties"`
	Triggers      string `xml:"triggers"`
	Parameters    string `xml:"parameters"`
	Options       string `xml:"options"`
}

func NewJobConfigDeclarativeJobPropertyTrackerAction() *JobConfigDeclarativeJobPropertyTrackerAction {
	return &JobConfigDeclarativeJobPropertyTrackerAction{}
}

func (action *JobConfigDeclarativeJobPropertyTrackerAction) GetId() string {
	return action.Id
}
