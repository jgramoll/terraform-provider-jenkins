package client

import (
	"encoding/xml"
	"fmt"
)

// JenkinsError Error response from jenkins
type JenkinsError struct {
	XMLName xml.Name `xml:"html"`
	Title   string
	Message string
}

// For error interface
func (r *JenkinsError) Error() string {
	return fmt.Sprintf("%v: %v", r.Title, r.Message)
}

func (jenkinsError *JenkinsError) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tok xml.Token
	var err error
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			switch elem.Name.Local {
			case "title":
				tok, err = d.Token()
				if err != nil {
					return err
				}
				if char, ok := tok.(xml.CharData); ok {
					jenkinsError.Title = string(char)
				}
			case "p":
				fallthrough
			case "pre":
				tok, err = d.Token()
				if err != nil {
					return err
				}
				if char, ok := tok.(xml.CharData); ok {
					jenkinsError.Message += string(char)
				}
			}
		}
	}
	return err
}
