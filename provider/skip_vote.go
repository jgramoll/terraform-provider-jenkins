package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type skipVote struct {
	OnSuccessful bool `mapstructure:"on_successful"`
	OnFailed     bool `mapstructure:"on_failed"`
	OnUnstable   bool `mapstructure:"on_unstable"`
	OnNotBuilt   bool `mapstructure:"on_not_built"`
}

func newSkipVote() *skipVote {
	return &skipVote{}
}

func newSkipVotefromClient(clientSkipVote *client.JobGerritTriggerSkipVote) *skipVote {
	newSkipVote := newSkipVote()
	newSkipVote.OnSuccessful = clientSkipVote.OnSuccessful
	newSkipVote.OnFailed = clientSkipVote.OnFailed
	newSkipVote.OnUnstable = clientSkipVote.OnUnstable
	newSkipVote.OnNotBuilt = clientSkipVote.OnNotBuilt
	return newSkipVote
}

func newClientSkipVote(v *[]*skipVote) *client.JobGerritTriggerSkipVote {
	clientSkipVote := client.NewJobGerritTriggerSkipVote()
	if v != nil && len(*v) != 0 {
		vote := (*v)[0]
		clientSkipVote.OnSuccessful = vote.OnSuccessful
		clientSkipVote.OnFailed = vote.OnFailed
		clientSkipVote.OnUnstable = vote.OnUnstable
		clientSkipVote.OnNotBuilt = vote.OnNotBuilt
	}
	return clientSkipVote
}
