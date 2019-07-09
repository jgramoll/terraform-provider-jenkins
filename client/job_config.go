package client

import "encoding/xml"

type jobConfig struct {
	XMLName xml.Name `xml:"flow-definition"`
	// actions
	Description string `xml:"description"`
	// keepDependencies
	// properties
	Definition *JobDefinition `xml:"definition"`
	Triggers   *[]*Trigger    `xml:"triggers"`
	Disabled   bool           `xml:"disabled"`
}

func JobConfigFromJob(job *Job) *jobConfig {
	return &jobConfig{}
}
