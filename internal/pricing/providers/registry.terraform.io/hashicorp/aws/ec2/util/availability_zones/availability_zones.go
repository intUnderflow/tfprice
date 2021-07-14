package availability_zones

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// GetAvailabilityZoneInfo gets metadata about an entire AWS availability zone
// Uses https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html
func GetAvailabilityZoneInfo(availabilityZone string) (*types.AvailabilityZone, error){
	client := ec2.New(ec2.Options{})

	zoneName := "zone-name"

	describeAZInput := ec2.DescribeAvailabilityZonesInput{
		Filters: []types.Filter{
			{
				Name:   &zoneName,
				Values: []string{availabilityZone},
			},
		},
	}

	response, err := client.DescribeAvailabilityZones(
		context.Background(),
		&describeAZInput,
	)
	if err != nil {
		return nil, err
	}

	if len(response.AvailabilityZones) != 1 {
		return nil, errors.New("could not locate exact availability zone")
	}

	return &response.AvailabilityZones[0], nil
}
