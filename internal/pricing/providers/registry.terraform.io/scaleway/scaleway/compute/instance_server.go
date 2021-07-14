package compute

import (
	"encoding/json"
	"errors"
	"fmt"
	scwtypes "github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/scaleway/scaleway/types"
	"github.com/intUnderflow/tfprice/internal/types"
	"time"
)

type Values struct {
	Type string  `json:"type"`
	Zone *string `json:"zone"`
}

func Price(resource types.Resource, config scwtypes.ProviderConfig) (*types.PriceRange, error) {
	// Extract resource values
	values := &Values{}
	err := json.Unmarshal(resource.Values, &values)
	if err != nil {
		return nil, err
	}

	// Calculate server zone
	// Server zone property in Values takes priority, otherwise fall back on the one set in
	// the ProviderConfig
	zone := config.Expressions.Zone.ConstantValue
	if values.Zone != nil {
		zone = values.Zone
	}

	if zone == nil {
		return nil, errors.New("cannot determine zone for server")
	}

	// Get the pricing for this Zone from the API
	zonePricing, err := GetServerPricingForZone(*zone)
	if err != nil {
		return nil, err
	}

	// TODO: Account for whether a flexible IP is assigned to this server or not (this causes a minor price difference)

	// Do we have info on the server?
	serverInfo := zonePricing.Servers[values.Type]

	if serverInfo == nil {
		return nil, errors.New(fmt.Sprintf(
			"cannot find pricing information for server type %s in zone %s",
			values.Type,
			*zone,
		))
	}

	prices := make([]types.PricePoint, 1)
	prices[0] = types.PricePoint{
		Duration: time.Hour,
		Cost: types.CurrencyAmount{
			Cost:         serverInfo.HourlyPrice * 100,
			CurrencyCode: "EUR",
		},
	}

	return &types.PriceRange{
		PricePoints: prices,
	}, nil
}
