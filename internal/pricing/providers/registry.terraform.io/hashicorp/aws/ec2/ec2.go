package ec2

import (
	"encoding/json"
	ec2types "github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/hashicorp/aws/ec2/types"
	awstypes "github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/hashicorp/aws/types"
	"github.com/intUnderflow/tfprice/internal/types"
)

func Price(resource types.Resource, config awstypes.AWSProviderConfig) (*types.PriceRange, error) {
	// Parse out resource values into an EC2 resource values object
	ec2Resource := ec2types.AWSEC2Resource{}
	err := json.Unmarshal(resource.Values, &ec2Resource)
	if err != nil {
		return nil, err
	}

	// TODO: Handle reserved capacity etc
	return onDemandPrice(ec2Resource, config)
}
