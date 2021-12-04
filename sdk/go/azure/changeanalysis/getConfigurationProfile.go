


package changeanalysis

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfile(ctx *pulumi.Context, args *LookupConfigurationProfileArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfileResult, error) {
	var rv LookupConfigurationProfileResult
	err := ctx.Invoke("azure-native:changeanalysis:getConfigurationProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfileArgs struct {
	ProfileName string `pulumi:"profileName"`
}


type LookupConfigurationProfileResult struct {
	Id         string                                         `pulumi:"id"`
	Identity   *ResourceIdentityResponse                      `pulumi:"identity"`
	Location   *string                                        `pulumi:"location"`
	Name       string                                         `pulumi:"name"`
	Properties ConfigurationProfileResourcePropertiesResponse `pulumi:"properties"`
	SystemData *SystemDataResponse                            `pulumi:"systemData"`
	Type       string                                         `pulumi:"type"`
}
