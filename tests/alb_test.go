package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestALB(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/alb",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	albDNSName := terraform.Output(t, terraformOptions, "dns_name")
	assert.NotEmpty(t, albDNSName)
}
