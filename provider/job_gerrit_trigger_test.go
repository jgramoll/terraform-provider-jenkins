package provider

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/hashicorp/terraform/helper/acctest"
// 	"github.com/hashicorp/terraform/helper/resource"
// 	"github.com/hashicorp/terraform/terraform"
// 	"github.com/jgramoll/terraform-provider-jenkins/client"
// )

// // func init() {
// // 	stageTypes["spinnaker_pipeline_destroy_server_group_stage"] = client.DestroyServerGroupStageType
// // }

// func TestAccJobGerritTriggerBasic(t *testing.T) {
// 	var jobRef client.Job
// 	// var stages []client.Stage
// 	jobName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
// 	// target := "my-target"
// 	// newTarget := "new-my-target"
// 	jobResourceName := "jenkins_job.test"
// 	trigger1 := "jenkins_job_gerrit_trigger.1"
// 	trigger2 := "jenkins_job_gerrit_trigger.2"

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		// CheckDestroy: testAccCheckJobTriggerDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccJobGerritTriggerConfigBasic(jobName, target, 2),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					// resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
// 					// resource.TestCheckResourceAttr(stage1, "target", target),
// 					// resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
// 					// resource.TestCheckResourceAttr(stage2, "target", target),
// 					testAccCheckJobExists(jobResourceName, &jobRef),
// 					testAccCheckJobTriggers(jobResourceName, []string{
// 						trigger1,
// 						trigger2,
// 					}, &triggers),
// 				),
// 			},
// 			{
// 				ResourceName: trigger1,
// 				ImportState:  true,
// 				ImportStateIdFunc: func(*terraform.State) (string, error) {
// 					if len(triggers) == 0 {
// 						return "", fmt.Errorf("no stages to import")
// 					}
// 					return fmt.Sprintf("%s_%s", jobRef.ID, stages[0].GetRefID()), nil
// 				},
// 				ImportStateVerify: true,
// 			},
// 			{
// 				ResourceName: trigger2,
// 				ImportState:  true,
// 				ImportStateIdFunc: func(*terraform.State) (string, error) {
// 					if len(stages) < 2 {
// 						return "", fmt.Errorf("no stages to import")
// 					}
// 					return fmt.Sprintf("%s_%s", jobRef.ID, stages[1].GetRefID()), nil
// 				},
// 				ImportStateVerify: true,
// 			},
// 			{
// 				Config: testAccJobGerritTriggerConfigBasic(pipeName, newTarget, 2),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					// resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
// 					// resource.TestCheckResourceAttr(stage1, "target", newTarget),
// 					// resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
// 					// resource.TestCheckResourceAttr(stage2, "target", newTarget),
// 					testAccCheckJobExists(jobResourceName, &jobRef),
// 					testAccCheckJobTriggers(jobResourceName, []string{
// 						trigger1,
// 						trigger2,
// 					}, &triggers),
// 				),
// 			},
// 			{
// 				Config: testAccJobGerritTriggerConfigBasic(pipeName, target, 1),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					// resource.TestCheckResourceAttr(trigger1, "name", "Stage 1"),
// 					// resource.TestCheckResourceAttr(trigger1, "target", target),
// 					testAccCheckJobExists(jobResourceName, &jobRef),
// 					testAccCheckJobStages(jobResourceName, []string{
// 						trigger1,
// 					}, &triggers),
// 				),
// 			},
// 			{
// 				Config: testAccJobGerritTriggerConfigBasic(pipeName, target, 0),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					testAccCheckJobExists(jobResourceName, &jobRef),
// 					testAccCheckJobTriggers(jobResourceName, []string{}, &triggers),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccJobGerritTriggerConfigBasic(pipeName string, target string, count int) string {
// 	triggers := ""
// 	for i := 1; i <= count; i++ {
// 		triggers += fmt.Sprintf(`
// resource "jenkins_job_gerrit_trigger" "%v" {
// 	job      = "${jenkins_job.test.id}"
// }`, i, i, target)
// 	}

// 	// name     = "Stage %v"
// 	// cluster  = "test_cluster"
// 	// target   = "%v"

// 	return fmt.Sprintf(`
// resource "jenkins_job" "test" {
// 	name        = "%s"
// }`, pipeName) + triggers
// }
