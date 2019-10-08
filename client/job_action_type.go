package client

import (
	"errors"
	"fmt"
)

type JobActionType string

var DeclarativeJobActionType JobActionType = "DeclarativeJobAction"
var DeclarativeJobPropertyTrackerActionType JobActionType = "DeclarativeJobPropertyTrackerAction"

func ParseJobActionType(s string) (JobActionType, error) {
	switch s {
	default:
		return "", errors.New(fmt.Sprintf("Unknown Action Type %s", s))
	case string(DeclarativeJobActionType):
		return DeclarativeJobActionType, nil
	case string(DeclarativeJobPropertyTrackerActionType):
		return DeclarativeJobPropertyTrackerActionType, nil
	}
}
