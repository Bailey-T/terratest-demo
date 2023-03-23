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
					/*
					BackendConfig: map[string]interface{}{
						"storage_account_name": os.Getenv("StorageAccount"),
						"key":                  os.Getenv("TF_VAR_uuid"),
						"container_name":       "terratest",
						"access_key":           os.Getenv("BackendAccessKey"),
					},
					*/
			})

			// Save options for later test stages
			test_structure.SaveTerraformOptions(t, fixtureFolder, terraformOptions)

			// Triggers the terraform init and terraform apply command
			terraform.InitAndApplyAndIdempotent(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "validate", func() {
			// run validation checks here
			//terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)
			//publicIpAddress := terraform.Output(t, terraformOptions, "public_ip_address")
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