package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraform(t *testing.T) {
    // Define the terraform options with default retryable errors to ignore
    terraformOptions := &terraform.Options{
        TerraformDir: "../",

        Vars: map[string]interface{}{
            "aws_region": "eu-central-1",
            "cluster_name": "eks-cluster",
        },

    // Clean up resources at the end of the test
    defer terraform.Destroy(t, terraformOptions)

    // Run 'terraform init' and 'terraform apply'. Fail the test if there are any errors.
    terraform.InitAndApply(t, terraformOptions)

    // Validate your code works as expected
    // Add assertions here to verify your infrastructure
}
