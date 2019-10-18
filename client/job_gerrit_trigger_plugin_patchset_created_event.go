package client

import "encoding/xml"

func init() {
	jobGerritTriggerOnEventsUnmarshalFunc["com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginPatchsetCreatedEvent"] = unmarshalJobGerritTriggerPluginPatchsetCreatedEvent
}

type JobGerritTriggerPluginPatchsetCreatedEvent struct {
	XMLName xml.Name `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginPatchsetCreatedEvent"`

	ExcludeDrafts        bool `xml:"excludeDrafts"`
	ExcludeTrivialRebase bool `xml:"excludeTrivialRebase"`
	ExcludeNoCodeChange  bool `xml:"excludeNoCodeChange"`
	ExcludePrivateState  bool `xml:"excludePrivateState"`
	ExcludeWipState      bool `xml:"excludeWipState"`
}

func NewJobGerritTriggerPluginPatchsetCreatedEvent() *JobGerritTriggerPluginPatchsetCreatedEvent {
	return &JobGerritTriggerPluginPatchsetCreatedEvent{
		ExcludeDrafts:        false,
		ExcludeTrivialRebase: false,
		ExcludeNoCodeChange:  false,
		ExcludePrivateState:  false,
		ExcludeWipState:      false,
	}
}

func (event *JobGerritTriggerPluginPatchsetCreatedEvent) GetType() JobGerritTriggerOnEventType {
	return PluginPatchsetCreatedEventType
}

func unmarshalJobGerritTriggerPluginPatchsetCreatedEvent(d *xml.Decoder, start xml.StartElement) (JobGerritTriggerOnEvent, error) {
	event := NewJobGerritTriggerPluginPatchsetCreatedEvent()
	err := d.DecodeElement(event, &start)
	if err != nil {
		return nil, err
	}
	return event, nil
}
