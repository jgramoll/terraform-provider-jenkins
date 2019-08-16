package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobParameterDefinition interface {
	fromClientJobParameterDefintion(client.JobParameterDefinition) (jobParameterDefinition, error)
	toClientJobParameterDefinition(id string) client.JobParameterDefinition
	setResourceData(*schema.ResourceData) error
}
