package provider

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGerritTriggerBasic(t *testing.T) {
	var jobRef client.Job
	var triggers []client.JobTrigger
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	trigger1 := "jenkins_job_gerrit_trigger.trigger_1"
	trigger2 := "jenkins_job_gerrit_trigger.trigger_2"
	serverName := "my-server"
	newServerName := "my-new-server"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritTriggerConfigServerName(jobName, serverName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "plugin", "my-plugin"),
					resource.TestCheckResourceAttr(trigger1, "server_name", serverName),
					resource.TestCheckResourceAttr(trigger1, "silent_mode", "false"),
					resource.TestCheckResourceAttr(trigger2, "plugin", "my-plugin"),
					resource.TestCheckResourceAttr(trigger2, "server_name", serverName),
					resource.TestCheckResourceAttr(trigger2, "silent_mode", "false"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobTriggers(&jobRef, []string{
						trigger1,
						trigger2,
					}, &triggers, ensureJobGerritTrigger),
				),
			},
			{
				Config: testAccJobGerritTriggerConfigServerName(jobName, newServerName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "plugin", "my-plugin"),
					resource.TestCheckResourceAttr(trigger1, "server_name", newServerName),
					resource.TestCheckResourceAttr(trigger1, "silent_mode", "false"),
					resource.TestCheckResourceAttr(trigger2, "plugin", "my-plugin"),
					resource.TestCheckResourceAttr(trigger2, "server_name", newServerName),
					resource.TestCheckResourceAttr(trigger2, "silent_mode", "false"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobTriggers(&jobRef, []string{
						trigger1,
						trigger2,
					}, &triggers, ensureJobGerritTrigger),
				),
			},
			{
				ResourceName:  trigger1,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Invalid trigger id"),
			},
			{
				ResourceName: trigger1,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(triggers) == 0 {
						return "", fmt.Errorf("no triggers to import")
					}
					propertyId := (*jobRef.Properties.Items)[0].GetId()
					return strings.Join([]string{jobName, propertyId, triggers[0].GetId()}, IdDelimiter), nil
				},
				ImportStateVerify: true,
			},
			{
				ResourceName: trigger2,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(triggers) == 0 {
						return "", fmt.Errorf("no triggers to import")
					}
					propertyId := (*jobRef.Properties.Items)[0].GetId()
					return strings.Join([]string{jobName, propertyId, triggers[1].GetId()}, IdDelimiter), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccJobGerritTriggerConfigServerName(jobName, serverName, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "plugin", "my-plugin"),
					resource.TestCheckResourceAttr(trigger1, "server_name", serverName),
					resource.TestCheckResourceAttr(trigger1, "silent_mode", "false"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobTriggers(&jobRef, []string{
						trigger1,
					}, &triggers, ensureJobGerritTrigger),
				),
			},
			{
				Config: testAccJobGerritTriggerConfigServerName(jobName, serverName, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobTriggers(&jobRef, []string{}, &triggers, ensureJobGerritTrigger),
				),
			},
		},
	})
}

func testAccJobGerritTriggerConfigBasic(jobName string) string {
	return testAccJobGerritTriggerConfigServerName(jobName, "__ANY__", 1)
}

func testAccJobGerritTriggerConfigServerName(jobName string, serverName string, count int) string {
	triggers := ""
	for i := 1; i <= count; i++ {
		triggers += fmt.Sprintf(`
resource "jenkins_job_gerrit_trigger" "trigger_%v" {
	property = "${jenkins_job_pipeline_triggers_property.prop_1.id}"

	plugin = "my-plugin"
	server_name = "%v"
	skip_vote {}
}`, i, serverName)
	}

	t := testAccJobPipelineTriggersPropertyConfigBasic(jobName, 1) + triggers
	return t
}

func ensureJobGerritTrigger(clientTriggerInterface client.JobTrigger, rs *terraform.ResourceState) error {
	triggerInterface, err := newJobGerritTrigger().fromClientJobTrigger(clientTriggerInterface)
	if err != nil {
		return err
	}
	trigger := triggerInterface.(*jobGerritTrigger)

	err = testCompareResourceBool("JobGerritTrigger", "SilentMode", rs.Primary.Attributes["silent_mode"], trigger.SilentMode)
	if err != nil {
		return err
	}
	err = testCompareResourceBool("JobGerritTrigger", "SilentStartMode", rs.Primary.Attributes["silent_start_mode"], trigger.SilentStartMode)
	if err != nil {
		return err
	}
	err = testCompareResourceBool("JobGerritTrigger", "EscapeQuotes", rs.Primary.Attributes["escape_quotes"], trigger.EscapeQuotes)
	if err != nil {
		return err
	}
	if trigger.NameAndEmailParameterMode != rs.Primary.Attributes["name_and_email_parameter_mode"] {
		return fmt.Errorf("expected name_and_email_parameter_mode %s, got %s",
			rs.Primary.Attributes["name_and_email_parameter_mode"], trigger.NameAndEmailParameterMode)
	}
	if trigger.CommitMessageParameterMode != rs.Primary.Attributes["commit_message_parameter_mode"] {
		return fmt.Errorf("expected commit_message_parameter_mode %s, got %s",
			rs.Primary.Attributes["commit_message_parameter_mode"], trigger.CommitMessageParameterMode)
	}
	if trigger.ChangeSubjectParameterMode != rs.Primary.Attributes["change_subject_parameter_mode"] {
		return fmt.Errorf("expected change_subject_parameter_mode %s, got %s",
			rs.Primary.Attributes["change_subject_parameter_mode"], trigger.ChangeSubjectParameterMode)
	}
	if trigger.CommentTextParameterMode != rs.Primary.Attributes["comment_text_parameter_mode"] {
		return fmt.Errorf("expected comment_text_parameter_mode %s, got %s",
			rs.Primary.Attributes["comment_text_parameter_mode"], trigger.CommentTextParameterMode)
	}
	if trigger.ServerName != rs.Primary.Attributes["server_name"] {
		return fmt.Errorf("expected server_name %s, got %s",
			rs.Primary.Attributes["server_name"], trigger.ServerName)
	}
	err = testCompareResourceBool("JobGerritTrigger", "DynamicTriggerConfiguration", rs.Primary.Attributes["dynamic_trigger_configuration"], trigger.DynamicTriggerConfiguration)
	if err != nil {
		return err
	}

	return nil
}
