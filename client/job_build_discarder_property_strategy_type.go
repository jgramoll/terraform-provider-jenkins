package client

import (
	"errors"
	"fmt"
)

type JobBuildDiscarderPropertyStrategyType string

var LogRotatorType JobBuildDiscarderPropertyStrategyType = "LogRotator"

func ParseJobBuildDiscarderPropertyStrategyType(s string) (JobBuildDiscarderPropertyStrategyType, error) {
	switch s {
	default:
		return "", errors.New(fmt.Sprintf("Unknown Build Discarder Property Strategy Type %s", s))
	case string(LogRotatorType):
		return LogRotatorType, nil
	}
}
