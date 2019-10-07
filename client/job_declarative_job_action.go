package client

import "encoding/xml"

func init() {
	jobActionUnmarshalFunc["org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction"] = unmarshalDeclarativeJobAction
}

type JobDeclarativeJobAction struct {
	XMLName xml.Name `xml:"org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction"`
	Id      string   `xml:"id,attr,omitempty"`
	Plugin  string   `xml:"plugin,attr,omitempty"`
}

func NewJobDeclarativeJobAction() *JobDeclarativeJobAction {
	return &JobDeclarativeJobAction{}
}

func (action *JobDeclarativeJobAction) GetId() string {
	return action.Id
}

func (a *JobDeclarativeJobAction) SetId(id string) {
	a.Id = id
}

func (*JobDeclarativeJobAction) GetType() JobActionType {
	return DeclarativeJobAction
}

func unmarshalDeclarativeJobAction(d *xml.Decoder, start xml.StartElement) (JobAction, error) {
	action := NewJobDeclarativeJobAction()
	err := d.DecodeElement(action, &start)
	if err != nil {
		return nil, err
	}
	return action, nil
}
