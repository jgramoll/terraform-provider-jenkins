package client

import (
	"encoding/xml"
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
