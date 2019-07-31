package client

import "encoding/xml"

type JobConfigDeclarativeJobAction struct {
	XMLName xml.Name `xml:"org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction"`
	Id      string   `xml:"id,attr,omitempty"`
}

func NewJobConfigDeclarativeJobAction() *JobConfigDeclarativeJobAction {
	return &JobConfigDeclarativeJobAction{}
}

func (action *JobConfigDeclarativeJobAction) GetId() string {
	return action.Id
}
