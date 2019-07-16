package client

type JobGerritTriggerSkipVote struct {
	OnSuccessful bool `xml:"onSuccessful"`
	OnFailed     bool `xml:"onFailed"`
	OnUnstable   bool `xml:"onUnstable"`
	OnNotBuilt   bool `xml:"onNotBuilt"`
}

func NewJobGerritTriggerSkipVote() *JobGerritTriggerSkipVote {
	return &JobGerritTriggerSkipVote{
		OnSuccessful: false,
		OnFailed: false,
		OnUnstable: false,
		OnNotBuilt: false,
	}
}
