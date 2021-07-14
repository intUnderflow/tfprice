package pricing

import (
	"github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/hashicorp/aws"
	"github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/scaleway/scaleway"
	"github.com/intUnderflow/tfprice/internal/types"
)

var pricingProviders = map[string]types.Provider{
	"registry.terraform.io/scaleway/scaleway": 	scaleway.Provider{},
	"registry.terraform.io/hashicorp/aws":		aws.Provider{},
}
