package pricing

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"time"
)

const (
	// HourlyCostRateCode and UpfrontCostRateCode sourced from
	// https://www.sentiatechblog.com/using-the-ec2-price-list-api
	HourlyCostRateCode = "6YS6EN2CT7"
	UpfrontCostRateCode = "2TG2D8R56U"
)

func GetPricingWithFilters(filters []types.Filter) ([]AWSPricing, error){
	// Send request to AWS and get result
	client := pricing.New(pricing.Options{})

	formatVersion := "aws_v1"
	getProductsInput := pricing.GetProductsInput{
		Filters: filters,
		FormatVersion: &formatVersion,
		MaxResults: 100,
	}

	paginator := pricing.NewGetProductsPaginator(client, &getProductsInput, getPaginatorOptions)

	productList := make([]AWSPricing, 0)

	for {
		output, err := paginator.NextPage(context.Background())

		if err != nil {
			return nil, err
		}

		for _, product := range output.PriceList {
			pricingObject := AWSPricing{}
			err := json.Unmarshal([]byte(product), &pricingObject)
			if err != nil {
				return nil, err
			}
			productList = append(productList, pricingObject)
		}

		if !paginator.HasMorePages() {
			break
		}
	}

	return productList, nil
}

func getPaginatorOptions(options *pricing.GetProductsPaginatorOptions){
	options.Limit = 100
	options.StopOnDuplicateToken = true
}

type AWSPricing struct {
	Product         AWSPricingProduct `json:"product"`
	ServiceCode     string            `json:"serviceCode"`
	Terms           AWSPricingTerms   `json:"terms"`
	Version         string            `json:"version"`
	PublicationDate time.Time         `json:"publicationDate"`
}

type AWSPricingProduct struct {
	ProductFamily string `json:"productFamily"`
	Attributes map[string]string `json:"attributes"`
	Sku string `json:"sku"`
}

type AWSPricingTerms struct {
	OnDemand map[string]json.RawMessage `json:"OnDemand"`
}