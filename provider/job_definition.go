package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobDefinition interface {
	fromClientJobDefintion(client.JobDefinition) jobDefinition
	toClientDefinition(definitionId string) client.JobDefinition
	setResourceData(*schema.ResourceData) error
}
