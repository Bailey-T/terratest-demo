package test

import (
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var uniqueId string = random.UniqueId()

func TestEndToEndDeploymentScenario(t *testing.T) {
	t.Parallel()

	fixtureFolder := "../terraform/test"

	// Use Terratest to deploy the infrastructure
	test_structure.RunTestStage(t, "setup", func() {
		terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
			// Indicate the directory that contains the Terraform configuration to deploy
			TerraformDir:    fixtureFolder,
			TerraformBinary: "terraform",
			MaxRetries:      3,
			Vars: map[string]interface{}{
				"guid": uniqueId,
			},
		})

		// Save options for later test stages
		test_structure.SaveTerraformOptions(t, fixtureFolder, terraformOptions)

		// Triggers the terraform init and terraform apply command
		terraform.InitAndApplyAndIdempotent(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "validate", func() {
		// run validation checks here
		terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)
		module1output := terraform.Output(t, terraformOptions, "module1")
		module2output := terraform.Output(t, terraformOptions, "module2")
		if module1output != "foo" {
			assert.Equal(t, "foo", module1output)
		}
		if module2output != "bar" {
			assert.Equal(t, "bar", module2output)
		}
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
