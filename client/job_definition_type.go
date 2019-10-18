package client

import (
	"errors"
	"fmt"
)

type JobDefinitionType string

var CpsScmFlowDefinitionType JobDefinitionType = "CpsScmFlowDefinition"

func ParseJobDefinitionType(s string) (JobDefinitionType, error) {
	switch s {
	default:
		return "", errors.New(fmt.Sprintf("Unknown Definition Type %s", s))
	case "CpsScmFlowDefinition":
		return CpsScmFlowDefinitionType, nil
	}
}
