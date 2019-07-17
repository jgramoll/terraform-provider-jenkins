package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobBuildDiscarderPropertyStrategy interface {
	fromClientStrategy(client.JobPipelineBuildDiscarderPropertyStrategy) jobBuildDiscarderPropertyStrategy
	toClientStrategy(id string) client.JobPipelineBuildDiscarderPropertyStrategy
	setResourceData(*schema.ResourceData) error
}
