package client

import (
	"encoding/xml"
	"errors"
	"fmt"
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

func (t *ParameterMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	switch s {
	default:
		return errors.New(fmt.Sprintf("Unknown Parameter Mode %s", s))
	case "PLAIN":
		*t = ParameterModePlain
	case "BASE64":
		*t = ParameterModeBase64
	}
	return nil
}
