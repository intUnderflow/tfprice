package compute

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	PricingApi = "https://api.scaleway.com/instance/v1/zones/%s/products/servers"
)

type ApiResponse struct {
	Servers map[string]*ServerProduct `json:"servers"`
}

type ServerProduct struct {
	MonthlyPrice float64 `json:"monthly_price"`
	HourlyPrice  float64 `json:"hourly_price"`
}

func GetServerPricingForZone(zone string) (*ApiResponse, error) {
	response, err := http.Get(fmt.Sprintf(PricingApi, zone))
	if err != nil {
		return nil, err
	}

	target := ApiResponse{}
	err = json.NewDecoder(response.Body).Decode(&target)
	if err != nil {
		return nil, err
	}
	return &target, nil
}
