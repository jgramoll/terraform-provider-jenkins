package client

import "encoding/xml"

type JobGerritTrigger struct {
	XMLName xml.Name `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.GerritTrigger"`

	// spec
	Projects *JobGerritTriggerProjects `xml:"gerritProjects"`
	// dynamicGerritProjects
	SkipVote   *JobGerritTriggerSkipVote `xml:"skipVote"`
	SilentMode bool                      `xml:"silentMode"`
	// notificationLevel
	SilentStartMode bool `xml:"silentStartMode"`
	EscapeQuotes    bool `xml:"escapeQuotes"`
	// NameAndEmailParameterMode string `xml:"nameAndEmailParameterMode"`
	// dependencyJobsNames
	// <commitMessageParameterMode>BASE64</commitMessageParameterMode>
	// <changeSubjectParameterMode>PLAIN</changeSubjectParameterMode>
	// <commentTextParameterMode>BASE64</commentTextParameterMode>
	// <buildStartMessage/>
	// <buildFailureMessage/>
	// <buildSuccessfulMessage/>
	// <buildUnstableMessage/>
	// <buildNotBuiltMessage/>
	// <buildUnsuccessfulFilepath/>
	// <customUrl/>
	ServerName string `xml:"serverName"`
	// <triggerOnEvents class="linked-list">...</triggerOnEvents>
	DynamicTriggerConfiguration bool `xml:"dynamicTriggerConfiguration"`
	// <triggerConfigURL/>
	// <triggerInformationAction/>
}
