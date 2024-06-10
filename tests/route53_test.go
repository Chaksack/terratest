package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRoute53(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/route53",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	zoneID := terraform.Output(t, terraformOptions, "zone_id")
	assert.NotEmpty(t, zoneID)
}
