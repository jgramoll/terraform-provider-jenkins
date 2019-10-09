package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobBuildDiscarderPropertyStrategy interface {
	fromClientStrategy(client.JobBuildDiscarderPropertyStrategy) (jobBuildDiscarderPropertyStrategy, error)
	toClientStrategy() client.JobBuildDiscarderPropertyStrategy
}
