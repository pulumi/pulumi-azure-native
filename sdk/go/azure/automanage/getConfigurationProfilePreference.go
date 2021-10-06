


package automanage

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfilePreference(ctx *pulumi.Context, args *LookupConfigurationProfilePreferenceArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfilePreferenceResult, error) {
	var rv LookupConfigurationProfilePreferenceResult
	err := ctx.Invoke("azure-native:automanage:getConfigurationProfilePreference", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfilePreferenceArgs struct {
	ConfigurationProfilePreferenceName string `pulumi:"configurationProfilePreferenceName"`
	ResourceGroupName                  string `pulumi:"resourceGroupName"`
}


type LookupConfigurationProfilePreferenceResult struct {
	Id         string                                           `pulumi:"id"`
	Location   string                                           `pulumi:"location"`
	Name       string                                           `pulumi:"name"`
	Properties ConfigurationProfilePreferencePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                                `pulumi:"tags"`
	Type       string                                           `pulumi:"type"`
}
