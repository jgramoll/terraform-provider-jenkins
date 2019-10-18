package client

import (
	"encoding/xml"
)

type jobConfig struct {
	XMLName xml.Name `xml:"flow-definition"`
	Plugin  string   `xml:"plugin,attr,omitempty"`

	Actions          *JobActions       `xml:"actions"`
	Description      string            `xml:"description"`
	KeepDependencies bool              `xml:"keepDependencies"`
	Properties       *JobProperties    `xml:"properties"`
	Definition       *JobDefinitionXml `xml:"definition"`
	Triggers         string            `xml:"triggers"`
	Disabled         bool              `xml:"disabled"`
}

func JobConfigFromJob(job *Job) *jobConfig {
	return &jobConfig{
		Plugin:           job.Plugin,
		Actions:          job.Actions,
		Description:      job.Description,
		KeepDependencies: job.KeepDependencies,
		Disabled:         job.Disabled,
		Properties:       job.Properties,
		Definition: &JobDefinitionXml{
			Item: job.Definition,
		},
	}
}
