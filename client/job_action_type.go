package client

import (
	"errors"
	"fmt"
)

type JobActionType string

var DeclarativeJobAction JobActionType = "DeclarativeJobAction"
var DeclarativeJobPropertyTrackerAction JobActionType = "DeclarativeJobPropertyTrackerAction"

func ParseJobActionType(s string) (JobActionType, error) {
	switch s {
	default:
		return "", errors.New(fmt.Sprintf("Unknown Action Type %s", s))
	case "DeclarativeJobAction":
		return DeclarativeJobAction, nil
	case "DeclarativeJobPropertyTrackerAction":
		return DeclarativeJobPropertyTrackerAction, nil
	}
}
