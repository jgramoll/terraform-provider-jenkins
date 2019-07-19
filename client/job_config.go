package client

import (
	"encoding/xml"
)

type jobConfig struct {
	XMLName xml.Name `xml:"flow-definition"`
	Id      string   `xml:"id,attr,omitempty"`
	Plugin  string   `xml:"plugin,attr,omitempty"`

	// actions
	Description      string            `xml:"description"`
	KeepDependencies bool              `xml:"keepDependencies"`
	Properties       *JobProperties    `xml:"properties"`
	Definition       *JobDefinitionXml `xml:"definition"`
	// triggers
	Disabled bool `xml:"disabled"`
}

func JobConfigFromJob(job *Job) *jobConfig {
	return &jobConfig{
		Description:      job.Description,
		KeepDependencies: job.KeepDependencies,
		Disabled:         job.Disabled,
		Properties:       job.Properties,
		Definition: &JobDefinitionXml{
			Item: job.Definition,
		},
	}
}
