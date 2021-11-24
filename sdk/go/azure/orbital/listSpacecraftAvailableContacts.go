


package orbital

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSpacecraftAvailableContacts(ctx *pulumi.Context, args *ListSpacecraftAvailableContactsArgs, opts ...pulumi.InvokeOption) (*ListSpacecraftAvailableContactsResult, error) {
	var rv ListSpacecraftAvailableContactsResult
	err := ctx.Invoke("azure-native:orbital:listSpacecraftAvailableContacts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSpacecraftAvailableContactsArgs struct {
	ContactProfile    ResourceReference `pulumi:"contactProfile"`
	EndTime           string            `pulumi:"endTime"`
	GroundStationName string            `pulumi:"groundStationName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SpacecraftName    string            `pulumi:"spacecraftName"`
	StartTime         string            `pulumi:"startTime"`
}


type ListSpacecraftAvailableContactsResult struct {
	NextLink string                      `pulumi:"nextLink"`
	Value    []AvailableContactsResponse `pulumi:"value"`
}
