package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobBuildDiscarderPropertyStrategyLogRotator struct {
	DaysToKeep         int `mapstructure:"days_to_keep"`
	NumToKeep          int `mapstructure:"num_to_keep"`
	ArtifactDaysToKeep int `mapstructure:"artifact_days_to_keep"`
	ArtifactNumToKeep  int `mapstructure:"artifact_num_to_keep"`
}

func newJobBuildDiscarderPropertyStrategyLogRotator() *jobBuildDiscarderPropertyStrategyLogRotator {
	return &jobBuildDiscarderPropertyStrategyLogRotator{
		DaysToKeep:         -1,
		NumToKeep:          -1,
		ArtifactDaysToKeep: -1,
		ArtifactNumToKeep:  -1,
	}
}

func (strategy *jobBuildDiscarderPropertyStrategyLogRotator) toClientStrategy(id string) client.JobPipelineBuildDiscarderPropertyStrategy {
	clientStrategy := client.NewJobPipelineBuildDiscarderPropertyStrategyLogRotator()
	clientStrategy.Id = id
	clientStrategy.DaysToKeep = strategy.DaysToKeep
	clientStrategy.NumToKeep = strategy.NumToKeep
	clientStrategy.ArtifactDaysToKeep = strategy.ArtifactDaysToKeep
	clientStrategy.ArtifactNumToKeep = strategy.ArtifactNumToKeep
	return clientStrategy
}

func (s *jobBuildDiscarderPropertyStrategyLogRotator) fromClientStrategy(cs client.JobPipelineBuildDiscarderPropertyStrategy) jobBuildDiscarderPropertyStrategy {
	clientStrategy := cs.(*client.JobPipelineBuildDiscarderPropertyStrategyLogRotator)
	strategy := newJobBuildDiscarderPropertyStrategyLogRotator()
	strategy.DaysToKeep = clientStrategy.DaysToKeep
	strategy.NumToKeep = clientStrategy.NumToKeep
	strategy.ArtifactDaysToKeep = clientStrategy.ArtifactDaysToKeep
	strategy.ArtifactNumToKeep = clientStrategy.ArtifactNumToKeep
	return strategy
}

func (strategy *jobBuildDiscarderPropertyStrategyLogRotator) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("days_to_keep", strategy.DaysToKeep); err != nil {
		return err
	}
	if err := d.Set("num_to_keep", strategy.NumToKeep); err != nil {
		return err
	}
	if err := d.Set("artifact_days_to_keep", strategy.ArtifactDaysToKeep); err != nil {
		return err
	}
	return d.Set("artifact_num_to_keep", strategy.ArtifactNumToKeep)
}
