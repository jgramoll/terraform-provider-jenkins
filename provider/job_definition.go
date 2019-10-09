package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobDefinition interface {
	fromClientDefinition(client.JobDefinition) (jobDefinition, error)
	toClientDefinition() (client.JobDefinition, error)
}
