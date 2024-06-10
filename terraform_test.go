package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformResources(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// Path to your Terraform code
		TerraformDir: "../",
	}

	// Run `terraform init` and `terraform apply`
	terraform.InitAndApply(t, terraformOptions)

	// Clean up resources with `terraform destroy` at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Example: Check if an output variable is set correctly
	instanceID := terraform.Output(t, terraformOptions, "instance_id")
	assert.NotEmpty(t, instanceID, "Instance ID should not be empty")
}
