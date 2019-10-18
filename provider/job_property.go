package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobProperty interface {
	fromClientProperty(client.JobProperty) (jobProperty, error)
	toClientProperty() (client.JobProperty, error)
}
