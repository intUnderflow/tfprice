package ec2

import (
	ec2types "github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/hashicorp/aws/ec2/types"
	awstypes "github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/hashicorp/aws/types"
	"github.com/intUnderflow/tfprice/internal/types"
)

func onDemandPrice(resource ec2types.AWSEC2Resource, config awstypes.AWSProviderConfig) (*types.PriceRange, error) {
}

