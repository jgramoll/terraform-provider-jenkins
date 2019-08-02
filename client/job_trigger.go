package client

type JobTrigger interface {
	GetId() string
	SetId(string)
}
