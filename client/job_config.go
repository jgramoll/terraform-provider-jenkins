package client

import (
	"encoding/xml"
)

type jobConfig struct {
	XMLName xml.Name `xml:"flow-definition"`
	Id      string   `xml:"id,attr"`
	Plugin  string   `xml:"plugin,attr"`

	// actions
	Description string `xml:"description"`
	// keepDependencies
	Properties *JobProperties    `xml:"properties"`
	Definition *JobDefinitionXml `xml:"definition"`
	// Triggers   *[]*Trigger    `xml:"triggers"`
	Disabled bool `xml:"disabled"`
}

func JobConfigFromJob(job *Job) *jobConfig {
	return &jobConfig{
		Id:          job.Id,
		Description: job.Description,
		Disabled:    job.Disabled,
		Properties:  job.Properties,
		Definition: &JobDefinitionXml{
			Item: job.Definition,
		},
	}
}
