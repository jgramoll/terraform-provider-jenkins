package client

import "encoding/xml"

func init() {
	propertyUnmarshalFunc["jenkins.model.BuildDiscarderProperty"] = unmarshalBuildDiscarderProperty
}

type JobBuildDiscarderProperty struct {
	XMLName xml.Name `xml:"jenkins.model.BuildDiscarderProperty"`
	Id      string   `xml:"id,attr"`

	Strategy *JobBuildDiscarderPropertyStrategyXml `xml:"strategy"`
}

func NewJobBuildDiscarderProperty() *JobBuildDiscarderProperty {
	return &JobBuildDiscarderProperty{
		Strategy: NewJobBuildDiscarderPropertyStrategyXml(),
	}
}

func (p *JobBuildDiscarderProperty) GetId() string {
	return p.Id
}

func (p *JobBuildDiscarderProperty) SetId(id string) {
	p.Id = id
}

func unmarshalBuildDiscarderProperty(d *xml.Decoder, start xml.StartElement) (JobProperty, error) {
	property := NewJobBuildDiscarderProperty()
	err := d.DecodeElement(property, &start)
	if err != nil {
		return nil, err
	}
	return property, nil
}
