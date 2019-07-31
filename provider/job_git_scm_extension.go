package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmExtension interface {
	fromClientExtension(client.GitScmExtension) jobGitScmExtension
	toClientExtension(id string) (client.GitScmExtension, error)
	setResourceData(*schema.ResourceData) error
}
