package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEKS(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/eks",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	clusterName := terraform.Output(t, terraformOptions, "cluster_name")
	assert.NotEmpty(t, clusterName)
}
