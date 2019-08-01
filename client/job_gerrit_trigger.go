package client

import (
	"encoding/xml"
	"errors"
)

func init() {
	jobTriggerUnmarshalFunc["com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.GerritTrigger"] = unmarshalJobGerritTrigger
}

// ErrJobGerritTriggerProjectNotFound job gerrit trigger project not found
var ErrJobGerritTriggerProjectNotFound = errors.New("Could not find job gerrit trigger project")

// ErrJobGerritTriggerEventNotFound job gerrit trigger event not found
var ErrJobGerritTriggerEventNotFound = errors.New("Could not find job gerrit trigger event")

type JobGerritTrigger struct {
	XMLName xml.Name `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.GerritTrigger"`
	Id      string   `xml:"id,attr,omitempty"`
	Plugin  string   `xml:"plugin,attr,omitempty"`

	Spec                  string                    `xml:"spec"`
	Projects              *JobGerritTriggerProjects `xml:"gerritProjects"`
	dynamicGerritProjects *DynamicGerritProjects    `xml:"dynamicGerritProjects"`
	SkipVote              *JobGerritTriggerSkipVote `xml:"skipVote"`
	SilentMode            bool                      `xml:"silentMode"`
	// notificationLevel
	SilentStartMode           bool          `xml:"silentStartMode"`
	EscapeQuotes              bool          `xml:"escapeQuotes"`
	NameAndEmailParameterMode ParameterMode `xml:"nameAndEmailParameterMode"`
	// dependencyJobsNames
	CommitMessageParameterMode ParameterMode `xml:"commitMessageParameterMode"`
	ChangeSubjectParameterMode ParameterMode `xml:"changeSubjectParameterMode"`
	CommentTextParameterMode   ParameterMode `xml:"commentTextParameterMode"`
	// buildStartMessage
	// buildFailureMessage
	// buildSuccessfulMessage
	// buildUnstableMessage
	// buildNotBuiltMessage
	// buildUnsuccessfulFilepath
	// customUrl
	ServerName                  string                    `xml:"serverName"`
	TriggerOnEvents             *JobGerritTriggerOnEvents `xml:"triggerOnEvents"`
	DynamicTriggerConfiguration bool                      `xml:"dynamicTriggerConfiguration"`
	// triggerConfigURL
	// triggerInformationAction
}

func NewJobGerritTrigger() *JobGerritTrigger {
	return &JobGerritTrigger{
		Projects:                    NewJobGerritTriggerProjects(),
		SkipVote:                    NewJobGerritTriggerSkipVote(),
		SilentMode:                  false,
		SilentStartMode:             false,
		EscapeQuotes:                true,
		NameAndEmailParameterMode:   ParameterModePlain,
		CommitMessageParameterMode:  ParameterModeBase64,
		ChangeSubjectParameterMode:  ParameterModePlain,
		CommentTextParameterMode:    ParameterModeBase64,
		ServerName:                  "__ANY__",
		TriggerOnEvents:             NewJobGerritTriggerOnEvents(),
		DynamicTriggerConfiguration: false,
	}
}

func (trigger *JobGerritTrigger) GetId() string {
	return trigger.Id
}

func unmarshalJobGerritTrigger(d *xml.Decoder, start xml.StartElement) (JobTrigger, error) {
	trigger := NewJobGerritTrigger()
	err := d.DecodeElement(trigger, &start)
	if err != nil {
		return nil, err
	}
	return trigger, nil
}

func (trigger *JobGerritTrigger) GetProject(projectId string) (*JobGerritTriggerProject, error) {
	projects := *(trigger.Projects).Items
	for _, project := range projects {
		if project.Id == projectId {
			return project, nil
		}
	}
	return nil, ErrJobGerritTriggerProjectNotFound
}

func (trigger *JobGerritTrigger) UpdateProject(project *JobGerritTriggerProject) error {
	oldProject, err := trigger.GetProject(project.Id)
	if err != nil {
		return err
	}
	*oldProject = *project
	return nil
}

func (trigger *JobGerritTrigger) DeleteProject(projectId string) error {
	projects := *(trigger.Projects).Items
	for i, p := range projects {
		if p.Id == projectId {
			*trigger.Projects.Items = append(projects[:i], projects[i+1:]...)
			return nil
		}
	}
	return ErrJobGerritTriggerProjectNotFound
}

func (trigger *JobGerritTrigger) GetEvent(eventId string) (JobGerritTriggerOnEvent, error) {
	for _, event := range *trigger.TriggerOnEvents.Items {
		if event.GetId() == eventId {
			return event, nil
		}
	}
	return nil, ErrJobGerritTriggerEventNotFound
}

func (trigger *JobGerritTrigger) UpdateEvent(event JobGerritTriggerOnEvent) error {
	eventId := event.GetId()
	events := *trigger.TriggerOnEvents.Items
	for i, e := range events {
		if e.GetId() == eventId {
			events[i] = event
			return nil
		}
	}
	return ErrJobGerritTriggerEventNotFound
}

func (trigger *JobGerritTrigger) DeleteEvent(eventId string) error {
	events := *trigger.TriggerOnEvents.Items
	for i, e := range events {
		if e.GetId() == eventId {
			*trigger.TriggerOnEvents.Items = append(events[:i], events[i+1:]...)
			return nil
		}
	}
	return ErrJobGerritTriggerEventNotFound
}
