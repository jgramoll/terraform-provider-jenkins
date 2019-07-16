package provider

// "github.com/hashicorp/terraform/helper/schema"

type jobTrigger interface {
	// fromClientStage(client.Stage) stage
	// toClientStage(*client.Config) (client.Stage, error)
	// SetResourceData(*schema.ResourceData) error
	// SetRefID(string)
	// GetRefID() string
}

// TODO why does this not like mapstructure
// type baseStage struct {
// 	Name  string           `mapstructure:"name"`
// 	RefID string           `mapstructure:"ref_id"`
// 	Type  client.StageType `mapstructure:"type"`
// }
