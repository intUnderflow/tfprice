package ssm

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// getRegionLongName gets the long name of an AWS region
// For example the input "us-west-1" will output "US West (N. California)"
func getRegionLongName(region string) (*string, error){
	client := ssm.New(ssm.Options{})

	paramName := fmt.Sprintf("/aws/service/global-infrastructure/regions/%s/longName", region)

	getParamInput := ssm.GetParameterInput{
		Name: &paramName,
	}

	response, err := client.GetParameter(context.Background(), &getParamInput)

	if err != nil  {
		return nil, err
	}

	return response.Parameter.Value, nil
}