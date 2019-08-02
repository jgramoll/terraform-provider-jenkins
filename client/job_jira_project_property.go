package client

import (
	"encoding/xml"
)

func init() {
	propertyUnmarshalFunc["hudson.plugins.jira.JiraProjectProperty"] = unmarshalJobJiraProjectProperty
}

type JobJiraProjectProperty struct {
	XMLName xml.Name `xml:"hudson.plugins.jira.JiraProjectProperty"`
	Id      string   `xml:"id,attr,omitempty"`
	Plugin  string   `xml:"plugin,attr,omitempty"`
}

func NewJobJiraProjectProperty() *JobJiraProjectProperty {
	return &JobJiraProjectProperty{}
}

func (property *JobJiraProjectProperty) GetId() string {
	return property.Id
}

func (p *JobJiraProjectProperty) SetId(id string) {
	p.Id = id
}

func unmarshalJobJiraProjectProperty(d *xml.Decoder, start xml.StartElement) (JobProperty, error) {
	property := NewJobJiraProjectProperty()
	err := d.DecodeElement(property, &start)
	if err != nil {
		return nil, err
	}
	return property, nil
}
