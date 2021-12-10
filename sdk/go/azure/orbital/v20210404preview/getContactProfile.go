


package v20210404preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContactProfile(ctx *pulumi.Context, args *LookupContactProfileArgs, opts ...pulumi.InvokeOption) (*LookupContactProfileResult, error) {
	var rv LookupContactProfileResult
	err := ctx.Invoke("azure-native:orbital/v20210404preview:getContactProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContactProfileArgs struct {
	ContactProfileName string `pulumi:"contactProfileName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupContactProfileResult struct {
	AutoTrackingConfiguration    *string                      `pulumi:"autoTrackingConfiguration"`
	Etag                         string                       `pulumi:"etag"`
	EventHubUri                  *string                      `pulumi:"eventHubUri"`
	Id                           string                       `pulumi:"id"`
	Links                        []ContactProfileLinkResponse `pulumi:"links"`
	Location                     string                       `pulumi:"location"`
	MinimumElevationDegrees      *float64                     `pulumi:"minimumElevationDegrees"`
	MinimumViableContactDuration *string                      `pulumi:"minimumViableContactDuration"`
	Name                         string                       `pulumi:"name"`
	SystemData                   SystemDataResponse           `pulumi:"systemData"`
	Tags                         map[string]string            `pulumi:"tags"`
	Type                         string                       `pulumi:"type"`
}
