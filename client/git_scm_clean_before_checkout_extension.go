package client

import "encoding/xml"

type GitScmCleanBeforeCheckoutExtension struct {
	XMLName xml.Name `xml:"hudson.plugins.git.extensions.impl.CleanBeforeCheckout"`
	Id      string   `xml:"id,attr"`
}

func NewGitScmCleanBeforeCheckoutExtension() *GitScmCleanBeforeCheckoutExtension {
	return &GitScmCleanBeforeCheckoutExtension{}
}

func (e *GitScmCleanBeforeCheckoutExtension) GetId() string {
	return e.Id
}
