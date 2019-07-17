package client

// Job Property
type JobParameterDefinitionProperty interface{}

type JobParameterDefinitionPropertyChoice struct {
	// DefaultParameterValue map[string]string `xml:"defaultParameterValue"`
	Description string `xml:"description"`
	Name        string `xml:"name"`
	// Type                  string            `xml:"type"`
	Choices *[]string `xml:"choices"`
}
