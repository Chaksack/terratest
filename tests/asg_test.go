package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestASG(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/asg",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	asgName := terraform.Output(t, terraformOptions, "asg_name")
	assert.NotEmpty(t, asgName)
}
