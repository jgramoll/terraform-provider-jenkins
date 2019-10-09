package client

import (
	"errors"
	"fmt"
)

type JobPropertyType string

var BuildDiscarderPropertyType JobPropertyType = "BuildDiscarderProperty"
var DatadogJobPropertyType JobPropertyType = "DatadogJobProperty"
var DisableConcurrentBuildsJobPropertyType JobPropertyType = "DisableConcurrentBuildsJobProperty"
var JiraProjectPropertyType JobPropertyType = "JiraProjectProperty"
var ParametersDefinitionPropertyType JobPropertyType = "ParametersDefinitionProperty"
var PipelineTriggersJobPropertyType JobPropertyType = "PipelineTriggersJobProperty"

func ParseJobPropertyType(s string) (JobPropertyType, error) {
	switch s {
	default:
		return "", errors.New(fmt.Sprintf("Unknown Property Type %s", s))
	case string(BuildDiscarderPropertyType):
		return BuildDiscarderPropertyType, nil
	case string(DatadogJobPropertyType):
		return DatadogJobPropertyType, nil
	case string(DisableConcurrentBuildsJobPropertyType):
		return DisableConcurrentBuildsJobPropertyType, nil
	case string(JiraProjectPropertyType):
		return JiraProjectPropertyType, nil
	case string(ParametersDefinitionPropertyType):
		return ParametersDefinitionPropertyType, nil
	case string(PipelineTriggersJobPropertyType):
		return PipelineTriggersJobPropertyType, nil
	}
}
