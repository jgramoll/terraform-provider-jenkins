package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobTrigger interface {
	fromClientJobTrigger(client.JobTrigger) (jobTrigger, error)
	toClientJobTrigger(id string) (client.JobTrigger, error)
	setResourceData(*schema.ResourceData) error
}
