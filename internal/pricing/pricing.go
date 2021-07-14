package pricing

import (
	"errors"
	"github.com/intUnderflow/tfprice/internal/types"
)

func PricePlan(plan types.TerraformPlan) (*types.PricedPlan, error) {
	pricedModules := make([]types.PricedModule, len(plan.PlannedValuesModules))
	i := 0
	for moduleName, module := range plan.PlannedValuesModules {
		pricedResources := make([]types.PricedResource, 0)
		for _, resource := range module.Resources {
			priceRange, err := getPriceOfResource(resource, plan.Configuration)
			if err != nil {
				return nil, err
			}
			if priceRange != nil {
				pricedResource := types.PricedResource{
					Address:    resource.Address,
					PriceRange: *priceRange,
				}
				pricedResources = append(pricedResources, pricedResource)
			}
		}
		pricedModules[i] = types.PricedModule{
			ModuleName:      moduleName,
			PricedResources: pricedResources,
		}
		i = i + 1
	}
	return &types.PricedPlan{
		PricedModules: pricedModules,
	}, nil
}

func getPriceOfResource(resource types.Resource, config types.PlanConfiguration) (*types.PriceRange, error) {
	provider := pricingProviders[resource.ProviderName]
	if provider != nil {
		return provider.Price(resource, config)
	}
	return nil, errors.New("cannot find provider " + resource.ProviderName)
}
