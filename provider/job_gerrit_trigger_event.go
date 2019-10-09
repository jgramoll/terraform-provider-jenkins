package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerEvent interface {
	fromClientTriggerEvent(client.JobGerritTriggerOnEvent) (jobGerritTriggerEvent, error)
	toClientTriggerEvent() (client.JobGerritTriggerOnEvent, error)
}
