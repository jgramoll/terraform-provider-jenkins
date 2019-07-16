package client

import "encoding/xml"

type JobGerritTriggerPluginPatchsetCreatedEvent struct {
	XMLName xml.Name `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginPatchsetCreatedEvent"`

	ExcludeDrafts        bool `xml:"excludeDrafts"`
	ExcludeTrivialRebase bool `xml:"excludeTrivialRebase"`
	ExcludeNoCodeChange  bool `xml:"excludeNoCodeChange"`
	ExcludePrivateState  bool `xml:"excludePrivateState"`
	ExcludeWipState      bool `xml:"excludeWipState"`
}
