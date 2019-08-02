package client

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
