package client

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type CompareType int

const (
	CompareTypeUnknown CompareType = iota
	CompareTypePlain
	CompareTypeRegExp
	CompareTypeAnt
)

func (t CompareType) String() string {
	return [...]string{"UNKNOWN", "PLAIN", "REG_EXP", "ANT"}[t]
}

func ParseCompareType(s string) (CompareType, error) {
	switch s {
	default:
		return CompareTypeUnknown, errors.New(fmt.Sprintf("Unknown Compare Type %s", s))
	case "PLAIN":
		return CompareTypePlain, nil
	case "REG_EXP":
		return CompareTypeRegExp, nil
	case "ANT":
		return CompareTypeAnt, nil
	}
}

func (t CompareType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

func (t *CompareType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	compareType, err := ParseCompareType(s)
	if err != nil {
		return err
	}
	*t = compareType
	return nil
}
