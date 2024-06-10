package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestApiGateway(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/api_gateway",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	apiGatewayID := terraform.Output(t, terraformOptions, "api_gateway_id")
	assert.NotEmpty(t, apiGatewayID)
}
