package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerSkipVotes []*jobGerritTriggerSkipVote

func (votes *jobGerritTriggerSkipVotes) toClientSkipVote() *client.JobGerritTriggerSkipVote {
	for _, vote := range *votes {
		return vote.toClientSkipVote()
	}
	return client.NewJobGerritTriggerSkipVote()
}

func (*jobGerritTriggerSkipVotes) fromClientSkipVote(clientSkipVote *client.JobGerritTriggerSkipVote) *jobGerritTriggerSkipVotes {
	skipVote := newSkipVotefromClient(clientSkipVote)
	return &jobGerritTriggerSkipVotes{skipVote}
}
