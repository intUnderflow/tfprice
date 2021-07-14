package types

import (
	"time"
)

type Provider interface {
	Price(resource Resource, config PlanConfiguration) (*PriceRange, error)
}

type PricedPlan struct {
	PricedModules []PricedModule
}

type PricedModule struct {
	ModuleName      string
	PricedResources []PricedResource
}

type PricedResource struct {
	Address    string
	PriceRange PriceRange
}

type PriceRange struct {
	PricePoints []PricePoint
}

type PricePoint struct {
	Duration time.Duration // For example if billed per hour this is 1 hour long
	Cost     CurrencyAmount
}

type CurrencyAmount struct {
	Cost         float64 // Cost is in lowest denominator of a currency, for example in cents for USD
	CurrencyCode string  // USD
}

// ToDuration is only accurate to the second!
func (p PricePoint) ToDuration(duration time.Duration) PricePoint {
	factor := duration.Seconds() / p.Duration.Seconds()
	return PricePoint{
		Duration: duration,
		Cost: CurrencyAmount{
			Cost:         p.Cost.Cost * factor,
			CurrencyCode: p.Cost.CurrencyCode,
		},
	}
}
