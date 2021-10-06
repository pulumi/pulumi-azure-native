


package v20181001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUserSettingsWithLocation(ctx *pulumi.Context, args *LookupUserSettingsWithLocationArgs, opts ...pulumi.InvokeOption) (*LookupUserSettingsWithLocationResult, error) {
	var rv LookupUserSettingsWithLocationResult
	err := ctx.Invoke("azure-native:portal/v20181001:getUserSettingsWithLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserSettingsWithLocationArgs struct {
	Location         string `pulumi:"location"`
	UserSettingsName string `pulumi:"userSettingsName"`
}


type LookupUserSettingsWithLocationResult struct {
	Properties UserPropertiesResponse `pulumi:"properties"`
}
