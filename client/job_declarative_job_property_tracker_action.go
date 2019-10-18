package client

import "encoding/xml"

func init() {
	jobActionUnmarshalFunc["org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction"] = unmarshalDeclarativeJobPropertyTrackerAction
}

type JobDeclarativeJobPropertyTrackerAction struct {
	XMLName xml.Name `xml:"org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction"`
	Plugin  string   `xml:"plugin,attr,omitempty"`

	JobProperties string `xml:"jobProperties"`
	Triggers      string `xml:"triggers"`
	Parameters    string `xml:"parameters"`
	Options       string `xml:"options"`
}

func NewJobDeclarativeJobPropertyTrackerAction() *JobDeclarativeJobPropertyTrackerAction {
	return &JobDeclarativeJobPropertyTrackerAction{}
}

func (*JobDeclarativeJobPropertyTrackerAction) GetType() JobActionType {
	return DeclarativeJobPropertyTrackerActionType
}

func unmarshalDeclarativeJobPropertyTrackerAction(d *xml.Decoder, start xml.StartElement) (JobAction, error) {
	action := NewJobDeclarativeJobPropertyTrackerAction()
	err := d.DecodeElement(action, &start)
	if err != nil {
		return nil, err
	}
	return action, nil
}
