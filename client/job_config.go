package client

// JobConfig
type JobConfig struct {
	// actions
	Description string `xml:"description"`
	// keepDependencies
	// properties
	Definition *JobDefinition `xml:"definition"`
	Triggers   *[]*Trigger    `xml:"triggers"`
	Disabled   bool           `xml:"disabled"`
}

// NewJobConfig return JobConfig object with default values
func NewJobConfig() *JobConfig {
	return &JobConfig{}
}
