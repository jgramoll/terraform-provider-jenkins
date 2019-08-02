package client

import "encoding/xml"

func init() {
	jobDefinitionUnmarshalFunc["org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition"] = unmarshalCpsScmFlowDefinition
}

type CpsScmFlowDefinition struct {
	Class  string `xml:"class,attr"`
	Id     string `xml:"id,attr,omitempty"`
	Plugin string `xml:"plugin,attr,omitempty"`

	SCM         *GitSCM `xml:"scm"`
	ScriptPath  string  `xml:"scriptPath"`
	Lightweight bool    `xml:"lightweight"`
}

func NewCpsScmFlowDefinition() *CpsScmFlowDefinition {
	return &CpsScmFlowDefinition{
		Class: "org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition",
	}
}

func (d *CpsScmFlowDefinition) GetId() string {
	return d.Id
}

func (d *CpsScmFlowDefinition) SetId(id string) {
	d.Id = id
}

func unmarshalCpsScmFlowDefinition(d *xml.Decoder, start xml.StartElement) (JobDefinition, error) {
	definition := NewCpsScmFlowDefinition()
	err := d.DecodeElement(definition, &start)
	if err != nil {
		return nil, err
	}
	return definition, nil
}
