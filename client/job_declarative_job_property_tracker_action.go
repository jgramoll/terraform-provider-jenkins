package client

import "encoding/xml"

type JobDeclarativeJobPropertyTrackerAction struct {
	XMLName xml.Name `xml:"org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction"`
	Id      string   `xml:"id,attr,omitempty"`

	JobProperties string `xml:"jobProperties"`
	Triggers      string `xml:"triggers"`
	Parameters    string `xml:"parameters"`
	Options       string `xml:"options"`
}

func NewJobDeclarativeJobPropertyTrackerAction() *JobDeclarativeJobPropertyTrackerAction {
	return &JobDeclarativeJobPropertyTrackerAction{}
}

func (action *JobDeclarativeJobPropertyTrackerAction) GetId() string {
	return action.Id
}
