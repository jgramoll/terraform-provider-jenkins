package client

type JobAction interface {
	GetId() string
	SetId(string)
	GetType() JobActionType
}
