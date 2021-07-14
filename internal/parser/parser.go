package parser

import (
	"encoding/json"
	"github.com/intUnderflow/tfprice/internal/types"
)

func ParseJSONPlan(JSONPlan []byte) (*types.TerraformPlan, error) {
	internalPlan := terraformPlanInternal{}
	err := json.Unmarshal(JSONPlan, &internalPlan)
	if err != nil {
		return nil, err
	}

	moduleMap := map[string]types.PlannedValuesModule{}
	for moduleName, moduleValues := range internalPlan.PlannedValues {
		parsedModule := types.PlannedValuesModule{}
		err := json.Unmarshal(moduleValues, &parsedModule)
		if err != nil {
			return nil, err
		}
		moduleMap[moduleName] = parsedModule
	}

	plan := types.TerraformPlan{
		PlannedValuesModules: moduleMap,
		Configuration:        internalPlan.Configuration,
	}

	return &plan, nil
}

type terraformPlanInternal struct {
	PlannedValues map[string]json.RawMessage `json:"planned_values"`
	Configuration types.PlanConfiguration    `json:"configuration"`
}
