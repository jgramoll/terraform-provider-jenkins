package client

import (
	"errors"
	"fmt"
)

type JobGerritTriggerOnEventType string

var PluginChangeMergedEventType JobGerritTriggerOnEventType = "PluginChangeMergedEvent"
var PluginDraftPublishedEventType JobGerritTriggerOnEventType = "PluginDraftPublishedEvent"
var PluginPatchsetCreatedEventType JobGerritTriggerOnEventType = "PluginPatchsetCreatedEvent"

func ParseJobGerritTriggerOnEventType(s string) (JobGerritTriggerOnEventType, error) {
	switch s {
	default:
		return "", errors.New(fmt.Sprintf("Unknown Gerrit Trigger Type %s", s))
	case string(PluginChangeMergedEventType):
		return PluginChangeMergedEventType, nil
	case string(PluginDraftPublishedEventType):
		return PluginDraftPublishedEventType, nil
	case string(PluginPatchsetCreatedEventType):
		return PluginPatchsetCreatedEventType, nil
	}
}
