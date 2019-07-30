package client

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type ParameterMode int

const (
	ParameterModeUnknown ParameterMode = iota
	ParameterModePlain
	ParameterModeBase64
)

func (t ParameterMode) String() string {
	return [...]string{"UNKNOWN", "PLAIN", "BASE64"}[t]
}

func ParseParameterMode(s string) (ParameterMode, error) {
	switch s {
	default:
		return ParameterModeUnknown, errors.New(fmt.Sprintf("Unknown Parameter Mode %s", s))
	case "PLAIN":
		return ParameterModePlain, nil
	case "BASE64":
		return ParameterModeBase64, nil
	}
}

func (t ParameterMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

func (t *ParameterMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	mode, err := ParseParameterMode(s)
	if err != nil {
		return err
	}

	*t = mode
	return nil
}
