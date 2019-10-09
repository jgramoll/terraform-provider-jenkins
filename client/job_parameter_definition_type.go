package client

import (
	"errors"
	"fmt"
)

type JobParameterDefinitionType string

var ChoiceParameterDefinitionType JobParameterDefinitionType = "ChoiceParameterDefinition"

func ParseJobParameterDefinitionType(s string) (JobParameterDefinitionType, error) {
	switch s {
	default:
		return "", errors.New(fmt.Sprintf("Unknown Parameter Definition Type %s", s))
	case string(ChoiceParameterDefinitionType):
		return ChoiceParameterDefinitionType, nil
	}
}
