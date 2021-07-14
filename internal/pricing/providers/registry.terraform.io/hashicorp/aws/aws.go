package aws

import (
	"encoding/json"
	"errors"
	"github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/hashicorp/aws/ec2"
	awstypes "github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/hashicorp/aws/types"
	"github.com/intUnderflow/tfprice/internal/types"
)

type Provider struct {}

func (p Provider) Price(
	resource types.Resource, config types.PlanConfiguration,
) (*types.PriceRange, error) {
	value, exists := config.ProviderConfig["AWS"]
	if !exists {
		return nil, errors.New("no provider configuration for AWS found in the Terraform plan")
	}

	awsConfig := awstypes.AWSProviderConfig{}

	err := json.Unmarshal(value, &awsConfig)
	if err != nil {
		return nil, err
	}

	if resource.Type == "aws_instance" {
		return ec2.Price(resource, awsConfig)
	}

	// If we get here we don't support pricing for this resource, so return nothing
	return nil, nil
}