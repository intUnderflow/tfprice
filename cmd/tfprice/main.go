package main

import (
	"flag"
	"fmt"
	"github.com/intUnderflow/tfprice/internal/cost_strategy"
	"github.com/intUnderflow/tfprice/internal/parser"
	"github.com/intUnderflow/tfprice/internal/pricing"
	"github.com/intUnderflow/tfprice/internal/types"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	var planFilePath string
	var planCostStrategy string

	flag.StringVar(&planFilePath, "f", "./plan.json", "path to terraform plan file")
	flag.StringVar(&planCostStrategy, "s", "average", "whether to report the first price result received, best case, average case or worst case cost possibilities")
	flag.Parse()

	if planCostStrategy != "average" && planCostStrategy != "best" && planCostStrategy != "worst" && planCostStrategy != "first" {
		log.Fatal("Invalid option for cost strategy, valid values are: first,average,best,worst")
	}

	log.Printf("Reading file from [%s].", planFilePath)

	jsonContents, err := ioutil.ReadFile(planFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	parsedPlan, err := parser.ParseJSONPlan(jsonContents)
	if err != nil {
		log.Fatal(err)
		return
	}

	pricedPlan, err := pricing.PricePlan(*parsedPlan)
	if err != nil {
		log.Fatal(err)
		return
	}

	pricesPerCurrencyTable := map[string]PriceCurrencyTable{}

	for _, module := range pricedPlan.PricedModules {
		for _, resource := range module.PricedResources {
			// Attempt to extract a price
			pricePointSelected, err := cost_strategy.GetPriceWithStrategy(
				resource.PriceRange,
				planCostStrategy,
			)
			if err != nil {
				log.Fatal(err)
				return
			}
			if pricePointSelected != nil {
				// If we got a price insert it into the table for the relevant currency
				currencyCode := pricePointSelected.Cost.CurrencyCode

				priceTable := pricesPerCurrencyTable[currencyCode]
				priceTable.Entries = append(priceTable.Entries, PriceCurrencyTableEntry{
					Module:     module.ModuleName,
					Address:    resource.Address,
					PricePoint: *pricePointSelected,
				})
				pricesPerCurrencyTable[currencyCode] = priceTable
			}
		}
	}

	// Calculate total price per hour
	currencyCostPerHour := map[string]float64{}
	for currency, priceTable := range pricesPerCurrencyTable {
		var total float64
		// For each entry convert to 1 hour and then add to the total for this currency
		for i := range priceTable.Entries {
			entry := priceTable.Entries[i]
			total += entry.PricePoint.ToDuration(time.Hour).Cost.Cost
		}
		currencyCostPerHour[currency] = total
	}

	// Display total price per hour
	var currencyCostPerHourStrings []string
	for currency := range currencyCostPerHour {
		currencyCostPerHourStrings = append(
			currencyCostPerHourStrings,
			fmt.Sprintf("%.2f %s", currencyCostPerHour[currency], currency),
		)
	}

	fmt.Printf("Cost per hour: %s", strings.Join(currencyCostPerHourStrings[:], " and "))
}

type PriceCurrencyTable struct {
	Entries []PriceCurrencyTableEntry
}

type PriceCurrencyTableEntry struct {
	Module     string
	Address    string
	PricePoint types.PricePoint
}
