package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerSkipVote struct {
	OnSuccessful bool `mapstructure:"on_successful"`
	OnFailed     bool `mapstructure:"on_failed"`
	OnUnstable   bool `mapstructure:"on_unstable"`
	OnNotBuilt   bool `mapstructure:"on_not_built"`
}

func newJobGerritTriggerSkipVote() *jobGerritTriggerSkipVote {
	return &jobGerritTriggerSkipVote{}
}

func newSkipVotefromClient(clientSkipVote *client.JobGerritTriggerSkipVote) *jobGerritTriggerSkipVote {
	newSkipVote := newJobGerritTriggerSkipVote()
	newSkipVote.OnSuccessful = clientSkipVote.OnSuccessful
	newSkipVote.OnFailed = clientSkipVote.OnFailed
	newSkipVote.OnUnstable = clientSkipVote.OnUnstable
	newSkipVote.OnNotBuilt = clientSkipVote.OnNotBuilt
	return newSkipVote
}

func (vote *jobGerritTriggerSkipVote) toClientSkipVote() *client.JobGerritTriggerSkipVote {
	clientSkipVote := client.NewJobGerritTriggerSkipVote()
	clientSkipVote.OnSuccessful = vote.OnSuccessful
	clientSkipVote.OnFailed = vote.OnFailed
	clientSkipVote.OnUnstable = vote.OnUnstable
	clientSkipVote.OnNotBuilt = vote.OnNotBuilt
	return clientSkipVote
}
