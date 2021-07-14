package filters

import (
	awstypes "github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/intUnderflow/tfprice/internal/pricing/providers/registry.terraform.io/hashicorp/aws/ec2/util/availability_zones"
	"github.com/intUnderflow/tfprice/internal/types"
)

func createFiltersForEC2Resource(resource types.Resource) ([]awstypes.Filter, error){
	// TODO: nail down exact filters to use

	filters := []awstypes.Filter{
		makeFilter("servicecode", "AmazonEC2"),
		// Filter to availability zone and Region (pricing objects can exist in both!)
		makeFilter("location", ),
	}

	return nil, nil
}

func makeFilter(field string, value string) awstypes.Filter {
	return awstypes.Filter{
		Field: &field,
		Type: awstypes.FilterTypeTermMatch,
		Value: &value,
	}
}

// getLocationFilter gets the filters to constrain pricing lookups to the location the resource is being placed
func getLocationFilter(resource types.Resource) ([]awstypes.Filter, error){
	// To get the zone name, we need to find the AMI being deployed and look at what zone its in

	// Get the zone name

	// We need the region long name as this is used in pricing
	availabilityZoneInfo := availability_zones.GetAvailabilityZoneInfo()
}