package client

import (
	"encoding/xml"
)

type jobActionUnmarshaler func(*xml.Decoder, xml.StartElement) (JobAction, error)

var jobActionUnmarshalFunc = map[string]jobActionUnmarshaler{}

type JobActions struct {
	XMLName xml.Name     `xml:"actions"`
	Items   *[]JobAction `xml:",any"`
}

func NewJobActions() *JobActions {
	return &JobActions{
		Items: &[]JobAction{},
	}
}

func (actions *JobActions) Append(action JobAction) *JobActions {
	newActions := NewJobActions()
	if actions != nil && actions.Items != nil {
		*newActions.Items = append(*actions.Items, action)
	} else {
		*newActions.Items = []JobAction{action}
	}
	return newActions
}

func (actions *JobActions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tok xml.Token
	var err error
	*actions = *NewJobActions()
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			if unmarshalXML, ok := jobActionUnmarshalFunc[elem.Name.Local]; ok {
				action, err := unmarshalXML(d, elem)
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
