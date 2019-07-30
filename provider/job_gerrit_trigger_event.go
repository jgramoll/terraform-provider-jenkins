package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerEvent interface {
	fromClientJobTriggerEvent(client.JobGerritTriggerOnEvent) jobGerritTriggerEvent
	toClientJobTriggerEvent(id string) (client.JobGerritTriggerOnEvent, error)
	setResourceData(*schema.ResourceData) error
}
