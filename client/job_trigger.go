package client

type JobTrigger interface {
	GetType() JobTriggerType
}
