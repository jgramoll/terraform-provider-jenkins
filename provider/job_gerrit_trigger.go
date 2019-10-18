package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobTriggerInitFunc[client.GerritTriggerType] = func() jobTrigger {
		return newJobGerritTrigger()
	}
}

type jobGerritTrigger struct {
	Type client.JobTriggerType `mapstructure:"type"`

	Plugin                      string `mapstructure:"plugin"`
	ServerName                  string `mapstructure:"server_name"`
	SilentMode                  bool   `mapstructure:"silent_mode"`
	SilentStartMode             bool   `mapstructure:"silent_start_mode"`
	EscapeQuotes                bool   `mapstructure:"escape_quotes"`
	NameAndEmailParameterMode   string `mapstructure:"name_and_email_parameter_mode"`
	CommitMessageParameterMode  string `mapstructure:"commit_message_parameter_mode"`
	ChangeSubjectParameterMode  string `mapstructure:"change_subject_parameter_mode"`
	CommentTextParameterMode    string `mapstructure:"comment_text_parameter_mode"`
	DynamicTriggerConfiguration bool   `mapstructure:"dynamic_trigger_configuration"`

	SkipVote        *jobGerritTriggerSkipVotes `mapstructure:"skip_vote"`
	GerritProjects  *jobGerritProjects         `mapstructure:"gerrit_project"`
	TriggerOnEvents *interfaceJobTriggerEvents `mapstructure:"trigger_on_event"`
}

func newJobGerritTrigger() *jobGerritTrigger {
	return &jobGerritTrigger{
		Type: client.GerritTriggerType,

		EscapeQuotes:                true,
		NameAndEmailParameterMode:   "PLAIN",
		CommitMessageParameterMode:  "BASE64",
		ChangeSubjectParameterMode:  "PLAIN",
		CommentTextParameterMode:    "BASE64",
		DynamicTriggerConfiguration: false,
	}
}

func (t *jobGerritTrigger) fromClientTrigger(clientTriggerInterface client.JobTrigger) (jobTrigger, error) {
	clientTrigger, ok := clientTriggerInterface.(*client.JobGerritTrigger)
	if !ok {
		return nil, fmt.Errorf("Strategy is not of expected type, expected *client.JobGerritTrigger, actually %s",
			reflect.TypeOf(clientTriggerInterface).String())
	}

	trigger := newJobGerritTrigger()
	trigger.Plugin = clientTrigger.Plugin
	trigger.ServerName = clientTrigger.ServerName
	trigger.SilentMode = clientTrigger.SilentMode
	trigger.SilentStartMode = clientTrigger.SilentStartMode
	trigger.EscapeQuotes = clientTrigger.EscapeQuotes
	trigger.NameAndEmailParameterMode = clientTrigger.NameAndEmailParameterMode.String()
	trigger.CommitMessageParameterMode = clientTrigger.CommitMessageParameterMode.String()
	trigger.ChangeSubjectParameterMode = clientTrigger.ChangeSubjectParameterMode.String()
	trigger.CommentTextParameterMode = clientTrigger.CommentTextParameterMode.String()
	trigger.DynamicTriggerConfiguration = clientTrigger.DynamicTriggerConfiguration
	trigger.SkipVote = trigger.SkipVote.fromClientSkipVote(clientTrigger.SkipVote)
	trigger.GerritProjects = trigger.GerritProjects.fromClientProjects(clientTrigger.Projects)

	triggerEvents, err := trigger.TriggerOnEvents.fromClientTriggerEvents(clientTrigger.TriggerOnEvents)
	if err != nil {
		return nil, err
	}
	trigger.TriggerOnEvents = triggerEvents

	return trigger, nil
}

func (t *jobGerritTrigger) toClientTrigger() (client.JobTrigger, error) {
	clientTrigger := client.NewJobGerritTrigger()
	clientTrigger.Plugin = t.Plugin
	clientTrigger.ServerName = t.ServerName
	clientTrigger.SilentMode = t.SilentMode
	clientTrigger.SilentStartMode = t.SilentStartMode
	clientTrigger.EscapeQuotes = t.EscapeQuotes
	clientTrigger.DynamicTriggerConfiguration = t.DynamicTriggerConfiguration
	clientTrigger.SkipVote = t.SkipVote.toClientSkipVote()

	err := t.parseParameterMode(clientTrigger)
	if err != nil {
		return nil, err
	}

	projects, err := t.GerritProjects.toClientProjects()
	if err != nil {
		return nil, err
	}
	clientTrigger.Projects = projects

	triggerOnEvents, err := t.TriggerOnEvents.toClientTriggerEvents()
	if err != nil {
		return nil, err
	}
	clientTrigger.TriggerOnEvents = triggerOnEvents

	return clientTrigger, nil
}

func (t *jobGerritTrigger) parseParameterMode(clientTrigger *client.JobGerritTrigger) error {
	mode, err := client.ParseParameterMode(t.NameAndEmailParameterMode)
	if err != nil {
		return err
	}
	clientTrigger.NameAndEmailParameterMode = mode
	mode, err = client.ParseParameterMode(t.CommitMessageParameterMode)
	if err != nil {
		return err
	}
	clientTrigger.CommitMessageParameterMode = mode
	mode, err = client.ParseParameterMode(t.ChangeSubjectParameterMode)
	if err != nil {
		return err
	}
	clientTrigger.ChangeSubjectParameterMode = mode
	mode, err = client.ParseParameterMode(t.CommentTextParameterMode)
	if err != nil {
		return err
	}
	clientTrigger.CommentTextParameterMode = mode
	return nil
}
