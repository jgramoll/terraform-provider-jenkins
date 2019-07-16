package client

type CpsScmFlowDefinition struct {
	Class string `xml:"class,attr"`

	Id          string  `xml:"id,attr"`
	SCM         *GitSCM `xml:"scm"`
	ScriptPath  string  `xml:"scriptPath"`
	Lightweight bool    `xml:"lightweight"`
}

func NewCpsScmFlowDefinition() *CpsScmFlowDefinition {
	return &CpsScmFlowDefinition{
		Class: "org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition",
	}
}

// func (definition *CpsScmFlowDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// 	start.Name.Local = "asdf"
// 	e.
// 	return nil
// }
