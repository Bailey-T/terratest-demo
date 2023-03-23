package test
import (
	"testing"
	"time"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/retry"
	//"github.com/stretchr/testify/assert"
)
func TestTerraformExample(t *testing.T) {
	// retryable errors in terraform testing. Details on options:
	// https://github.com/gruntwork-io/terratest/blob/d1db0095a436b62ed92730ac12ec79497b47d2ee/modules/terraform/options.go#L40
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformBinary: "terraform",
		TerraformDir: "../terraform/",
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

	defer retry.DoWithRetry(t, "Terraform Destroy", 2, time.Duration(60), func() (string, error) {
		results, err := terraform.DestroyE(t, terraformOptions)
		return results, err
	})

	terraform.InitAndApplyAndIdempotent(t, terraformOptions)

	//output := terraform.Output(t, terraformOptions, "hello_world")
	//assert.Equal(t, "Hello, World!", output)
}