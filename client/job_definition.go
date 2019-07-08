package client

// JobDefinition
type JobDefinition struct {
	SCM         *SCM   `xml:"scm"`
	ScriptPath  string `xml:"scriptPath"`
	Lightweight bool   `xml:"lightweight"`
}
