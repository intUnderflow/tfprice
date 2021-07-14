package cost_strategy

import (
	"errors"
	"fmt"
	"github.com/intUnderflow/tfprice/internal/types"
	"time"
)

const (
	durationToUse = time.Second
)

// GetPriceWithStrategy gets a PricePoint after using a price strategy, returns this as a PricePoint
func GetPriceWithStrategy(priceRange types.PriceRange, strategy string) (*types.PricePoint, error) {
	if len(priceRange.PricePoints) == 0 {
		return nil, nil
	}
	switch strategy {
	case "first":
		return firstCase(priceRange)
	case "average":
		return averageCase(priceRange)
	default:
		return nil, errors.New(fmt.Sprintf("cost strategy %s not implemented", strategy))
	}
}

func firstCase(priceRange types.PriceRange) (*types.PricePoint, error) {
	return &priceRange.PricePoints[0], nil
}

func averageCase(priceRange types.PriceRange) (*types.PricePoint, error) {
	// Average up the total cost of all prices.
	// Right now we lock to the first currency seen and it is invalid to specify multiple currencies for the
	// same PriceRange.
	// TODO: Is there a better way to handle multiple prices in different currencies?
	var currencySeen *string
	currencySeen = nil

	sum := float64(0)
	sum = 0
	for _, pricePoint := range priceRange.PricePoints {
		// Validate currency is the same
		if currencySeen == nil {
			currencySeen = &pricePoint.Cost.CurrencyCode
		} else if currencySeen != &pricePoint.Cost.CurrencyCode {
			return nil, errors.New(
				fmt.Sprintf(
					"multiple currencies in a single PriceRange are not supported %s != %s",
					*currencySeen,
					pricePoint.Cost.CurrencyCode,
				),
			)
		}

		// To create the averages, we break everything down a common duration
		normalizedPoint := pricePoint.ToDuration(durationToUse)

		sum = sum + normalizedPoint.Cost.Cost
	}

	return &types.PricePoint{
		Cost: types.CurrencyAmount{
			Cost:         sum / float64(len(priceRange.PricePoints)),
			CurrencyCode: *currencySeen,
		},
		Duration: durationToUse,
	}, nil
}
