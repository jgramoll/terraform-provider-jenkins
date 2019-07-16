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
	NameAndEmailParameterMode ParameterMode `xml:"nameAndEmailParameterMode"`
	// dependencyJobsNames
	CommitMessageParameterMode ParameterMode `xml:"commitMessageParameterMode"`
	ChangeSubjectParameterMode ParameterMode `xml:"changeSubjectParameterMode"`
	CommentTextParameterMode ParameterMode `xml:"commentTextParameterMode"`
	// <buildStartMessage/>
	// <buildFailureMessage/>
	// <buildSuccessfulMessage/>
	// <buildUnstableMessage/>
	// <buildNotBuiltMessage/>
	// <buildUnsuccessfulFilepath/>
	// <customUrl/>
	ServerName string `xml:"serverName"`
	TriggerOnEvents *JobGerritTriggerOnEvents `xml:"triggerOnEvents"`
	DynamicTriggerConfiguration bool `xml:"dynamicTriggerConfiguration"`
	// <triggerConfigURL/>
	// <triggerInformationAction/>
}

func NewJobGerritTrigger() *JobGerritTrigger {
	return &JobGerritTrigger{
		Projects: NewJobGerritTriggerProjects(),
		SkipVote: NewJobGerritTriggerSkipVote(),
		SilentMode:false,
		SilentStartMode: false,
		EscapeQuotes: true,
		NameAndEmailParameterMode: ParameterModePlain,
		CommitMessageParameterMode: ParameterModeBase64,
		ChangeSubjectParameterMode: ParameterModePlain,
		CommentTextParameterMode: ParameterModeBase64,
		ServerName: "__ANY__",
		TriggerOnEvents: &JobGerritTriggerOnEvents{},
		DynamicTriggerConfiguration: false,
	}
}
