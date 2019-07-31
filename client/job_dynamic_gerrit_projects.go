package client

import (
	"encoding/xml"
)

type DynamicGerritProjects struct {
	XMLName xml.Name `xml:"dynamicGerritProjects"`
	Class   string   `xml:"class,attr"`
}

func NewDynamicGerritProjects() *DynamicGerritProjects {
	return &DynamicGerritProjects{
		Class: "empty-list",
	}
}
