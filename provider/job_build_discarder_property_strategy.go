package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobBuildDiscarderPropertyStrategy interface {
	fromClientStrategy(client.JobBuildDiscarderPropertyStrategy) jobBuildDiscarderPropertyStrategy
	toClientStrategy(id string) client.JobBuildDiscarderPropertyStrategy
	setResourceData(*schema.ResourceData) error
}
