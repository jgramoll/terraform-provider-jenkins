package provider

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

// Services used by provider
type Services struct {
	Config             client.Config
	JobService   client.JobService
}

// Config for provider
type Config struct {
	Address   string `mapstructure:"address"`
	Username  string `mapstructure:"username"`
	Token     string `mapstructure:"token"`
}

// Provider for terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_ADDRESS", nil),
				Description: "Address of jenkins",
			},

			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_USERNAME", nil),
				Description: "Name of the user to authenticate with jenkins",
			},

			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_TOKEN", nil),
				Description: "Token for the user to authenticate with jenkins",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"jenkins_job":              jobResource(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var config Config
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &config); err != nil {
		return nil, err
	}

	log.Println("[INFO] Initializing jenkins client")

	clientConfig := client.Config(config)
	c := client.NewClient(clientConfig)
	return &Services{
		Config:             clientConfig,
		JobService:    client.JobService{Client: c},
	}, nil
}
