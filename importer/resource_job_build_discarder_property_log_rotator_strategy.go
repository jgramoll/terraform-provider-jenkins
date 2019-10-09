package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobBuildDiscarderPropertyStrategyCodeFuncs["*client.JobBuildDiscarderPropertyStrategyLogRotator"] = jobBuildDiscarderPropertyLogRotatorStrategyCode
}

func jobBuildDiscarderPropertyLogRotatorStrategyCode(strategyInterface client.JobBuildDiscarderPropertyStrategy) string {
	strategy := strategyInterface.(*client.JobBuildDiscarderPropertyStrategyLogRotator)
	return fmt.Sprintf(`
    strategy {
      type  = "LogRotator"

      days_to_keep          = "%v"
      num_to_keep           = "%v"
      artifact_days_to_keep = "%v"
      artifact_num_to_keep  = "%v"
    }
`, strategy.DaysToKeep, strategy.NumToKeep,
		strategy.ArtifactDaysToKeep, strategy.ArtifactNumToKeep)
}
