package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobAction interface {
	fromClientAction(client.JobAction) (jobAction, error)
	toClientAction() (client.JobAction, error)
}
