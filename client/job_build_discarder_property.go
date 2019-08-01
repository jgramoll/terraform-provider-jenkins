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

func (property *JobBuildDiscarderProperty) GetId() string {
	return property.Id
}

func unmarshalBuildDiscarderProperty(d *xml.Decoder, start xml.StartElement) (JobProperty, error) {
	property := NewJobBuildDiscarderProperty()
	err := d.DecodeElement(property, &start)
	if err != nil {
		return nil, err
	}
	return property, nil
}
