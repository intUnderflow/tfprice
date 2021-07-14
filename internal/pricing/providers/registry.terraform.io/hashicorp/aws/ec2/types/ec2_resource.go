package types

type AWSEC2Resource struct {
	AvailabilityZone string `json:"availability_zone"`
	InstanceType string `json:"instance_type"`
}
