package client

import (
	"errors"
	"fmt"
)

type JobTriggerType string

var GerritTriggerType JobTriggerType = "GerritTrigger"

func ParseJobTriggerType(s string) (JobTriggerType, error) {
	switch s {
	default:
		return "", errors.New(fmt.Sprintf("Unknown Job Trigger Type %s", s))
	case string(GerritTriggerType):
		return GerritTriggerType, nil
	}
}
