package test

import (
	"testing"
	"time"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/retry"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	//"github.com/stretchr/testify/assert"
)
func TestEndToEndDeploymentScenario(t *testing.T) {
	t.Parallel()

	fixtureFolder := "../terraform/"

	// Use Terratest to deploy the infrastructure
	test_structure.RunTestStage(t, "setup", func() {
			terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
					// Indicate the directory that contains the Terraform configuration to deploy
					TerraformDir: fixtureFolder,
					TerraformBinary: "terraform",
					MaxRetries:   3,
			})

			// Save options for later test stages
			test_structure.SaveTerraformOptions(t, fixtureFolder, terraformOptions)

			// Triggers the terraform init and terraform apply command
			terraform.InitAndApplyAndIdempotent(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "validate", func() {
			// run validation checks here
			//terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)
			module1output := terraform.Output(t, terraformOptions, "module.thing1.someoutput")
	})

	// When the test is completed, teardown the infrastructure by calling terraform destroy
	test_structure.RunTestStage(t, "teardown", func() {
			terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)
			retry.DoWithRetry(t, "Terraform Destroy", 2, time.Duration(60), func() (string, error) {
				results, err := terraform.DestroyE(t, terraformOptions)
				return results, err
			})
	})
}