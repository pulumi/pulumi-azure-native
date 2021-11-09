


package automanage

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfile(ctx *pulumi.Context, args *LookupConfigurationProfileArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfileResult, error) {
	var rv LookupConfigurationProfileResult
	err := ctx.Invoke("azure-native:automanage:getConfigurationProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfileArgs struct {
	ConfigurationProfileName string `pulumi:"configurationProfileName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type LookupConfigurationProfileResult struct {
	Id         string                                 `pulumi:"id"`
	Location   string                                 `pulumi:"location"`
	Name       string                                 `pulumi:"name"`
	Properties ConfigurationProfilePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                     `pulumi:"systemData"`
	Tags       map[string]string                      `pulumi:"tags"`
	Type       string                                 `pulumi:"type"`
}
