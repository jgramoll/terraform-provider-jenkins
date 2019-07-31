package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobAction interface {
	fromClientAction(client.JobAction) (jobAction, error)
	toClientAction(id string) (client.JobAction, error)
	setResourceData(*schema.ResourceData) error
}
