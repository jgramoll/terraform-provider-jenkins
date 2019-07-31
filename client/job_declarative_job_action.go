package client

import "encoding/xml"

type JobDeclarativeJobAction struct {
	XMLName xml.Name `xml:"org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction"`
	Id      string   `xml:"id,attr,omitempty"`
}

func NewJobDeclarativeJobAction() *JobDeclarativeJobAction {
	return &JobDeclarativeJobAction{}
}

func (action *JobDeclarativeJobAction) GetId() string {
	return action.Id
}
