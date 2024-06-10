package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestElasticache(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/elasticcache",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	cacheClusterID := terraform.Output(t, terraformOptions, "cache_cluster_id")
	assert.NotEmpty(t, cacheClusterID)
}
