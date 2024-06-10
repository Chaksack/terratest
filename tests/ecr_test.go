package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestECR(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/ecr",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	repoURL := terraform.Output(t, terraformOptions, "repository_url")
	assert.NotEmpty(t, repoURL)
}
