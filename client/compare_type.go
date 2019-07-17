package client

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type CompareType int

const (
	CompareTypePlain CompareType = iota
	CompareTypeRegExp
)

func (t CompareType) String() string {
	return [...]string{"PLAIN", "REG_EXP"}[t]
}

func (t CompareType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

func (t *CompareType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	switch s {
	default:
		return errors.New(fmt.Sprintf("Unknown Compare Type %s", s))
	case "PLAIN":
		*t = CompareTypePlain
	case "REG_EXP":
		*t = CompareTypeRegExp
	}
	return nil
}
