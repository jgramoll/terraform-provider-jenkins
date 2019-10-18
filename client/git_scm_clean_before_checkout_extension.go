package client

import "encoding/xml"

func init() {
	scmExtensionUnmarshalFunc["hudson.plugins.git.extensions.impl.CleanBeforeCheckout"] = unmarshalScmExtension
}

type GitScmCleanBeforeCheckoutExtension struct {
	XMLName xml.Name `xml:"hudson.plugins.git.extensions.impl.CleanBeforeCheckout"`
}

func NewGitScmCleanBeforeCheckoutExtension() *GitScmCleanBeforeCheckoutExtension {
	return &GitScmCleanBeforeCheckoutExtension{}
}

func (e *GitScmCleanBeforeCheckoutExtension) GetType() GitScmExtensionType {
	return CleanBeforeCheckoutType
}

func unmarshalScmExtension(d *xml.Decoder, start xml.StartElement) (GitScmExtension, error) {
	extension := NewGitScmCleanBeforeCheckoutExtension()
	err := d.DecodeElement(extension, &start)
	if err != nil {
		return nil, err
	}
	return extension, nil
}
