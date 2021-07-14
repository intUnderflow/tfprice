package scaleway

import (
	"encoding/json"
	"github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/scaleway/scaleway/compute"
	scwtypes "github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/scaleway/scaleway/types"
	"github.com/intUnderflow/tfprice/internal/types"
)

type Provider struct{}

func (p Provider) Price(
	resource types.Resource, config types.PlanConfiguration,
) (*types.PriceRange, error) {
	// Parse out plan configuration for Scaleway
	scwConfig := scwtypes.ProviderConfig{}
	err := json.Unmarshal(config.ProviderConfig["scaleway"], &scwConfig)
	if err != nil {
		return nil, err
	}

	if resource.Type == "scaleway_instance_server" {
		return compute.Price(resource, scwConfig)
	}
	return nil, nil
}
