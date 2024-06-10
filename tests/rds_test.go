package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRDS(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/rds",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	dbInstanceEndpoint := terraform.Output(t, terraformOptions, "db_instance_endpoint")
	assert.NotEmpty(t, dbInstanceEndpoint)
}
