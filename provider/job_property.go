package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobProperty interface {
	fromClientJobProperty(client.JobProperty) (jobProperty, error)
	toClientProperty(id string) client.JobProperty
	setResourceData(*schema.ResourceData) error
}
