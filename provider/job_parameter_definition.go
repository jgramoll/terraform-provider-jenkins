package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobParameterDefinition interface {
	fromClientJobParameterDefinition(client.JobParameterDefinition) (jobParameterDefinition, error)
	toClientJobParameterDefinition() (client.JobParameterDefinition, error)
}
