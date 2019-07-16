package client

import (
	"encoding/xml"
)

type ParameterMode int

const (
	ParameterModePlain ParameterMode = iota
	ParameterModeBase64
)

func (t ParameterMode) String() string {
	return [...]string{"PLAIN", "BASE64"}[t]
}

func (t ParameterMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}
