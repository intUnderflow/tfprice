package types

import "encoding/json"

type TerraformPlan struct {
	PlannedValuesModules map[string]PlannedValuesModule
	Configuration        PlanConfiguration
}

type PlannedValuesModule struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Address      string          `json:"address"`
	ProviderName string          `json:"provider_name"`
	Type         string          `json:"type"`
	Values       json.RawMessage `json:"values"`
}

type PlanConfiguration struct {
	ProviderConfig map[string]json.RawMessage `json:"provider_config"`
}
