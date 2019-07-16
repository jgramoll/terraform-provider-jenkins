package client

type JobGerritTriggerSkipVote struct {
	OnSuccessful bool `xml:"onSuccessful"`
	OnFailed     bool `xml:"onFailed"`
	OnUnstable   bool `xml:"onUnstable"`
	OnNotBuilt   bool `xml:"onNotBuilt"`
}
