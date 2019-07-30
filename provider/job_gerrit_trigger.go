package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTrigger struct {
	ServerName                 string       `mapstructure:"server_name"`
	SilentMode                 bool         `mapstructure:"silent_mode"`
	SilentStartMode            bool         `mapstructure:"silent_start_mode"`
	EscapeQuotes               bool         `mapstructure:"escape_quotes"`
	NameAndEmailParameterMode  string       `mapstructure:"name_and_email_parameter_mode"`
	CommitMessageParameterMode string       `mapstructure:"commit_message_parameter_mode"`
	ChangeSubjectParameterMode string       `mapstructure:"change_subject_parameter_mode"`
	CommentTextParameterMode   string       `mapstructure:"comment_text_parameter_mode"`
	SkipVote                   *[]*skipVote `mapstructure:"skip_vote"`
}

func newJobGerritTrigger() *jobGerritTrigger {
	return &jobGerritTrigger{
		EscapeQuotes:               true,
		NameAndEmailParameterMode:  "PLAIN",
		CommitMessageParameterMode: "BASE64",
		ChangeSubjectParameterMode: "PLAIN",
		CommentTextParameterMode:   "BASE64",
		SkipVote:                   &[]*skipVote{},
	}
}

func (t *jobGerritTrigger) fromClientJobTrigger(clientTriggerInterface client.JobTrigger) jobTrigger {
	clientTrigger := clientTriggerInterface.(*client.JobGerritTrigger)
	trigger := newJobGerritTrigger()
	trigger.ServerName = clientTrigger.ServerName
	trigger.SilentMode = clientTrigger.SilentMode
	trigger.SilentStartMode = clientTrigger.SilentStartMode
	trigger.EscapeQuotes = clientTrigger.EscapeQuotes
	trigger.NameAndEmailParameterMode = clientTrigger.NameAndEmailParameterMode.String()
	trigger.CommitMessageParameterMode = clientTrigger.CommitMessageParameterMode.String()
	trigger.ChangeSubjectParameterMode = clientTrigger.ChangeSubjectParameterMode.String()
	trigger.CommentTextParameterMode = clientTrigger.CommentTextParameterMode.String()
	*trigger.SkipVote = []*skipVote{newSkipVotefromClient(clientTrigger.SkipVote)}
	return trigger
}

func (t *jobGerritTrigger) toClientJobTrigger(id string) (client.JobTrigger, error) {
	clientTrigger := client.NewJobGerritTrigger()
	clientTrigger.Id = id
	clientTrigger.ServerName = t.ServerName
	clientTrigger.SilentMode = t.SilentMode
	clientTrigger.SilentStartMode = t.SilentStartMode
	clientTrigger.EscapeQuotes = t.EscapeQuotes
	err := t.parseParameterMode(clientTrigger)
	if err != nil {
		return nil, err
	}
	clientTrigger.SkipVote = newClientSkipVote(t.SkipVote)
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

func (t *jobGerritTrigger) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("server_name", t.ServerName); err != nil {
		return err
	}
	if err := d.Set("silent_mode", t.SilentMode); err != nil {
		return err
	}
	if err := d.Set("silent_start_mode", t.SilentStartMode); err != nil {
		return err
	}
	if err := d.Set("escape_quotes", t.EscapeQuotes); err != nil {
		return err
	}
	if err := d.Set("name_and_email_parameter_mode", t.NameAndEmailParameterMode); err != nil {
		return err
	}
	if err := d.Set("commit_message_parameter_mode", t.CommitMessageParameterMode); err != nil {
		return err
	}
	if err := d.Set("change_subject_parameter_mode", t.ChangeSubjectParameterMode); err != nil {
		return err
	}
	if err := d.Set("comment_text_parameter_mode", t.CommentTextParameterMode); err != nil {
		return err
	}
	return d.Set("skip_vote", t.SkipVote)
}
