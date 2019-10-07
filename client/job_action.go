package client

type JobAction interface {
	GetType() JobActionType
}
