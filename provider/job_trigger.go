package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobTrigger interface {
	fromClientTrigger(client.JobTrigger) (jobTrigger, error)
	toClientTrigger() (client.JobTrigger, error)
}
