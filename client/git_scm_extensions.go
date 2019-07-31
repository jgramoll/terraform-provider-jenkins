package client

import (
	"encoding/xml"
)

type GitScmExtensions struct {
	Items *[]GitScmExtension
}

func NewGitScmExtensions() *GitScmExtensions {
	return &GitScmExtensions{
		Items: &[]GitScmExtension{},
	}
}

func (extensions *GitScmExtensions) Append(extension GitScmExtension) *GitScmExtensions {
	newExtensions := NewGitScmExtensions()
	if extensions.Items != nil {
		*newExtensions.Items = append(*extensions.Items, extension)
	} else {
		*newExtensions.Items = []GitScmExtension{extension}
	}
	return newExtensions
}

func (extensions *GitScmExtensions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tok xml.Token
	var err error
	extensions.Items = &[]GitScmExtension{}
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if elem, ok := tok.(xml.StartElement); ok {
			// TODO use map
			switch elem.Name.Local {
			case "hudson.plugins.git.extensions.impl.CleanBeforeCheckout":
				extension := NewGitScmCleanBeforeCheckoutExtension()
				err := d.DecodeElement(extension, &elem)
				if err != nil {
					return err
				}
				*extensions.Items = append(*extensions.Items, extension)
			}
		}
		if end, ok := tok.(xml.EndElement); ok {
			if end.Name.Local == "extensions" {
				break
			}
		}
	}
	return err
}
