package client

import "encoding/xml"

func init() {
	jobActionUnmarshalFunc["org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction"] = unmarshalDeclarativeJobPropertyTrackerAction
}

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

func unmarshalDeclarativeJobPropertyTrackerAction(d *xml.Decoder, start xml.StartElement) (JobAction, error) {
	action := NewJobDeclarativeJobPropertyTrackerAction()
	err := d.DecodeElement(action, &start)
	if err != nil {
		return nil, err
	}
	return action, nil
}
