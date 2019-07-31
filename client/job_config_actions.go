package client

import (
	"encoding/xml"
)

type JobConfigActions struct {
	XMLName xml.Name           `xml:"actions"`
	Items   *[]JobConfigAction `xml:",any"`
}

func NewJobConfigActions() *JobConfigActions {
	return &JobConfigActions{
		Items: &[]JobConfigAction{},
	}
}

func (actions *JobConfigActions) Append(action JobConfigAction) *JobConfigActions {
	newActions := NewJobConfigActions()
	if actions.Items != nil {
		*newActions.Items = append(*actions.Items, action)
	} else {
		*newActions.Items = []JobConfigAction{action}
	}
	return newActions
}

func (actions *JobConfigActions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tok xml.Token
	var err error
	actions.Items = &[]JobConfigAction{}
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			// TODO use map
			switch elem.Name.Local {
			case "org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction":
				action := NewJobConfigDeclarativeJobAction()
				err := d.DecodeElement(action, &elem)
				if err != nil {
					return err
				}
				*actions.Items = append(*actions.Items, action)
			case "org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction":
				action := NewJobConfigDeclarativeJobPropertyTrackerAction()
				err := d.DecodeElement(action, &elem)
				if err != nil {
					return err
				}
				*actions.Items = append(*actions.Items, action)
			}
		}
		if end, ok := tok.(xml.EndElement); ok {
			if end.Name.Local == "actions" {
				break
			}
		}
	}
	return err
}
