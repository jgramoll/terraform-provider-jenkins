package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobBuildDiscarderPropertyStrategyInitFunc[client.LogRotatorType] = func() jobBuildDiscarderPropertyStrategy {
		return newJobBuildDiscarderPropertyStrategyLogRotator()
	}
}

type jobBuildDiscarderPropertyStrategyLogRotator struct {
	Type client.JobBuildDiscarderPropertyStrategyType `mapstructure:"type"`

	DaysToKeep         int `mapstructure:"days_to_keep"`
	NumToKeep          int `mapstructure:"num_to_keep"`
	ArtifactDaysToKeep int `mapstructure:"artifact_days_to_keep"`
	ArtifactNumToKeep  int `mapstructure:"artifact_num_to_keep"`
}

func newJobBuildDiscarderPropertyStrategyLogRotator() *jobBuildDiscarderPropertyStrategyLogRotator {
	return &jobBuildDiscarderPropertyStrategyLogRotator{
		Type:               client.LogRotatorType,
		DaysToKeep:         -1,
		NumToKeep:          -1,
		ArtifactDaysToKeep: -1,
		ArtifactNumToKeep:  -1,
	}
}

func (strategy *jobBuildDiscarderPropertyStrategyLogRotator) toClientStrategy() client.JobBuildDiscarderPropertyStrategy {
	clientStrategy := client.NewJobBuildDiscarderPropertyStrategyLogRotator()
	clientStrategy.DaysToKeep = strategy.DaysToKeep
	clientStrategy.NumToKeep = strategy.NumToKeep
	clientStrategy.ArtifactDaysToKeep = strategy.ArtifactDaysToKeep
	clientStrategy.ArtifactNumToKeep = strategy.ArtifactNumToKeep
	return clientStrategy
}

func (*jobBuildDiscarderPropertyStrategyLogRotator) fromClientStrategy(clientStrategyInterface client.JobBuildDiscarderPropertyStrategy) (jobBuildDiscarderPropertyStrategy, error) {
	clientStrategy, ok := clientStrategyInterface.(*client.JobBuildDiscarderPropertyStrategyLogRotator)
	if !ok {
		return nil, fmt.Errorf("Strategy is not of expected type, expected *client.JobBuildDiscarderPropertyStrategyLogRotator, actually %s",
			reflect.TypeOf(clientStrategyInterface).String())
	}
	strategy := newJobBuildDiscarderPropertyStrategyLogRotator()
	strategy.DaysToKeep = clientStrategy.DaysToKeep
	strategy.NumToKeep = clientStrategy.NumToKeep
	strategy.ArtifactDaysToKeep = clientStrategy.ArtifactDaysToKeep
	strategy.ArtifactNumToKeep = clientStrategy.ArtifactNumToKeep
	return strategy, nil
}
